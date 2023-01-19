package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/mercadolibre/hsm-lib-poc/internal/hsm"
	"github.com/mercadolibre/hsm-lib-poc/internal/hsm/handler"
)

const (
	publicRSAKeyFormatDER = "30820122300D06092A864886F70D01010105000382010F003082010A0282010100AB7F4566FCA0F17FC3671E707B2235050B506119D2456AB89C229F9BC86C129E721E92D0B209DD2436E19DD5250C82A238FB654AAEBC444A46191AE0C14189040AFBB0230C4C7E0929E0BC1554887BD6B7B369F8CB9AFBB7720E59BA450848B886E7ED3F089F3360CC0DE683CAA8D578F2251658A1E9F418EE52CB40E257A3C7C89E75C7BA2E751DA915B5A0ED132C698574D3E893465838323EA62D3D98EB20EAEA663310B8DB553FD26F7BA8FBD88F590DACED25071B6924703A37F6C6AC2DA3BDC9A195125383C5407B680F881D93F3F333C1D10CC4F2272FC9803A3ED418DAEE438C717DB160F2A475D0CB0ADEE3C488A37E8C4B3984C041D50E84E0C0D50203010001"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	priKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil
	}

	return priKey, &priKey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(priKey *rsa.PrivateKey) string {
	priKeyBytes := x509.MarshalPKCS1PrivateKey(priKey)
	priKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: priKeyBytes,
		},
	)

	return string(priKeyPem)
}

func ParseRsaPrivateKeyFromPemStr(priPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(priPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pri, nil
}

func ExportRsaPublicKeyAsPemStr(pubKey *rsa.PublicKey) (string, error) {
	pubKeyBytes := x509.MarshalPKCS1PublicKey(pubKey)
	pubKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		},
	)

	return string(pubKeyPem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break
	}
	return nil, errors.New("key type is not RSA")
}

func ParseRSAPublicKeyFromPemStr(pubPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func formatRSAPublicKey(pubKey string) string {
	result := "-----BEGIN RSA PUBLIC KEY-----\n"
	iniPos := 0
	endPos := 64
	for {
		if endPos < len(pubKey) {
			result += pubKey[iniPos:endPos] + "\n"
			iniPos = endPos
			endPos += 64
			continue
		}
		result += pubKey[iniPos:] + "\n-----END RSA PUBLIC KEY-----"
		break
	}

	return result
}

func getRequestHashSum() string {
	msgHash := sha256.New()
	msgHash.Write([]byte("1234"))

	return fmt.Sprintf("%x\n", msgHash.Sum(nil))
}

func main() {
	hexMsg, err := hex.DecodeString(publicRSAKeyFormatDER)
	if err != nil {
		fmt.Println("Hex decode string fail ", err)
		return
	}

	pubKey := base64.StdEncoding.EncodeToString(hexMsg)
	pubKey = formatRSAPublicKey(pubKey)
	rsaPubKey, err := ParseRsaPublicKeyFromPemStr(pubKey)
	if err != nil {
		fmt.Println("Error parsing RSA public key from PEM string", err)
		return
	}

	fmt.Println(pubKey)
	// Encrypting a clear PIN
	pin := "1234"
	pin = hex.EncodeToString([]byte(pin))
	fmt.Println("PIN HEX:", pin)

	encBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPubKey, []byte("1234"), nil)
	if err != nil {
		fmt.Println("Error encrypting PIN 1 --> ", err)
		return
	}

	fmt.Printf("\n\nCiphertext PIN HEX: %x\n", encBytes)
	encBytes, err = rsa.EncryptOAEP(sha512.New(), rand.Reader, rsaPubKey, []byte(pin), nil)
	if err != nil {
		fmt.Println("Error encrypting PIN 2 --> ", err)
		return
	}

	fmt.Printf("\n\nCiphertext clear PIN: %x\n", encBytes)

	if run() != nil {
		fmt.Println("Fail running main")
	}
}

func run() error {
	// Start fury application
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	// Handler creation
	hsmService := hsm.NewHSMService()
	hsmHandler := handler.NewHSMHandler(hsmService)

	// HSM functionalities
	app.Post("/hsm/arqc-validation", hsmHandler.ARQCValidation)
	app.Post("/hsm/pin-generation", hsmHandler.PINGeneration)
	app.Post("/hsm/pvv-generation", hsmHandler.PVVGeneration)
	app.Post("/hsm/pin-block-generation", hsmHandler.PINBlockGeneration)
	app.Post("/hsm/pin-verification", hsmHandler.PINVerification)
	app.Post("/hsm/arpc-generation", hsmHandler.ARPCGeneration)
	app.Post("/hsm/generate-vd", hsmHandler.GenerateValidationData)
	app.Post("/hsm/validate-vv", hsmHandler.ValidateValidationData)
	return app.Run()
}
