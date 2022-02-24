package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	// 1.获得一个证书序列号
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 创建一个证书主题结构
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "localhost",
		//Country: nil,
		//Locality:           nil,
		//Province:           nil,
		//StreetAddress:      nil,
		//PostalCode:         nil,
		//SerialNumber:       "",
		//Names:              nil,
		//ExtraNames:         nil,
	}

	template := x509.Certificate{
		SerialNumber: serialNumber, // 证书序列号，用于记录由CA分发的唯一号码
		Subject:      subject,      // 证书的的主题
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 证书的有效期
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, // 该x509证书是用于服务器身份验证的
		IPAddresses:  []net.IP{net.ParseIP("192.168.24.147")},        // 该证书只能在该地址上运行
		//Raw:                         nil,
		//RawTBSCertificate:           nil,
		//RawSubjectPublicKeyInfo:     nil,
		//RawSubject:                  nil,
		//RawIssuer:                   nil,
		//Signature:                   nil,
		//SignatureAlgorithm:          0,
		//PublicKeyAlgorithm:          0,
		//PublicKey:                   nil,
		//Version:                     0,
		//Issuer:                      pkix.Name{},
		//Extensions:                  nil,
		//ExtraExtensions:             nil,
		//UnhandledCriticalExtensions: nil,
		//UnknownExtKeyUsage:          nil,
		//BasicConstraintsValid:       false,
		//IsCA:                  false,
		//MaxPathLen:            0,
		//MaxPathLenZero:        false,
		//SubjectKeyId:          nil,
		//AuthorityKeyId:        nil,
		//OCSPServer:            nil,
		//IssuingCertificateURL: nil,
		//DNSNames:              nil,
		//EmailAddresses:        nil,
		//URIs:                  nil,
		//PermittedDNSDomainsCritical: false,
		//PermittedDNSDomains:         nil,
		//ExcludedDNSDomains:          nil,
		//PermittedIPRanges:           nil,
		//ExcludedIPRanges:            nil,
		//PermittedEmailAddresses:     nil,
		//ExcludedEmailAddresses:      nil,
		//PermittedURIDomains:         nil,
		//ExcludedURIDomains:          nil,
		//CRLDistributionPoints:       nil,
		//PolicyIdentifiers:           nil,
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)                                                 // 生成一个RSA私钥（私钥结构里面包含一个公钥）
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) //创建一个经过DER编码格式化的切片

	certOut, _ := os.Create("server.crt")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}) // 以PEM 编码方式将SSL证书数据编码到cert.pem文件中
	certOut.Close()

	keyOut, _ := os.Create("server.key")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}) // 以PEM 编码方式将私钥数据编码到key.pem文件中
	keyOut.Close()
}
