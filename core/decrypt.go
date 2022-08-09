package core

func Decrypt(value string) (string, error) {
	decryptedByteArr, err := decryptToByte(value)
	if decryptedByteArr == nil {
		return "", err
	}
	return string(decryptedByteArr[:]), nil
}
