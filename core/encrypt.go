package core

func Encrypt(value string) (string, error) {
	encryptedByteArr, err := encryptToByte(value)
	if err != nil {
		return "", err
	}
	return encodeBase64(encryptedByteArr), nil
}
