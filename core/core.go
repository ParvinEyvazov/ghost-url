package core

import (
	"crypto/cipher"

	"encoding/base64"

	"golang.org/x/crypto/blowfish"
)

const _KEY string = "suPerHarD!To*FinD"

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func blowfishChecksizeAndPad(value []byte) []byte {
	modulus := len(value) % blowfish.BlockSize
	if modulus != 0 {
		padnglen := blowfish.BlockSize - modulus
		for i := 0; i < padnglen; i++ {
			value = append(value, 0)
		}
	}
	return value
}

func blowfishEncrypt(value, key []byte) ([]byte, error) {
	bcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	returnMe := make([]byte, blowfish.BlockSize+len(value))
	eiv := returnMe[:blowfish.BlockSize]
	ecbc := cipher.NewCBCEncrypter(bcipher, eiv)
	ecbc.CryptBlocks(returnMe[blowfish.BlockSize:], value)
	return returnMe, nil
}

func encryptToByte(value string) ([]byte, error) {
	var returnMe, valueInByteArr, paddedByteArr, keyByteArr []byte
	valueInByteArr = []byte(value)
	keyByteArr = []byte(_KEY)
	paddedByteArr = blowfishChecksizeAndPad(valueInByteArr)
	returnMe, err := blowfishEncrypt(paddedByteArr, keyByteArr)
	if err != nil {
		return nil, err
	}
	return returnMe, nil
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func blowfishDecrypt(value, key []byte) ([]byte, error) {
	dcipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	div := value[:blowfish.BlockSize]
	decrypted := value[blowfish.BlockSize:]
	if len(decrypted)%blowfish.BlockSize != 0 {
		return nil, err
	}
	dcbc := cipher.NewCBCDecrypter(dcipher, div)
	dcbc.CryptBlocks(decrypted, decrypted)
	return decrypted, nil
}

func decryptToByte(value string) ([]byte, error) {
	var returnMe, keyByteArr []byte
	keyByteArr = []byte(_KEY)
	decodeB64, err1 := decodeBase64(value)
	if err1 != nil {
		return nil, err1
	}
	returnMe, err2 := blowfishDecrypt(decodeB64, keyByteArr)
	if err2 != nil {
		return nil, err2
	}
	return returnMe, nil
}
