setup:
	mv static/wasm_exec.js.disabled static/wasm_exec.js

unsetup:
	mv static/wasm_exec.js static/wasm_exec.js.disabled

build:
	GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main.go

smallcss: 
	echo "go to https://www.textfixer.com/tools/remove-line-breaks.php"

run:
	go run ./cmd/webserver/main.go

runlive:
	nodemon --exec go run ./cmd/webserver/main.go --signal SIGTERM

br:
	make build
	make run


brlive:
	make build
	nodemon --exec go run ./cmd/webserver/main.go --signal SIGTERM

git:
	mv static/wasm_exec.js static/wasm_exec.js.disabled
	git add .
	git commit -m "$m"
	git push
