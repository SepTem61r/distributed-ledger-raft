package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"log"
)

// 生成ecdsa秘钥,返回私钥 账户地址
func GenerateKeyPair() (*ecdsa.PrivateKey, string) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panicf("ecdsa密钥生成失败: %v", err)
	}
	pubKeyBytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	address := hex.EncodeToString([]byte(pubKeyBytes))
	return privateKey, address
}

// sign 签名
func Sign(privateKey *ecdsa.PrivateKey, dataHash string) []byte {
	hashBytes, _ := hex.DecodeString(dataHash)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hashBytes)
	if err != nil {
		log.Panicf("签名失败：%v", err)
	}
	return signature
}

// 风险测评
func Verity(address string, dataHash string, signature []byte) bool {
	hashBytes, _ := hex.DecodeString(dataHash)
	pubKeyBytes, err := hex.DecodeString(address)
	if err != nil {
		return false
	}
	genericPubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return false
	}
	pubKey, ok := genericPubKey.(*ecdsa.PublicKey)
	if !ok {
		return false
	}
	return ecdsa.VerifyASN1(pubKey, hashBytes, signature)
}
