package helpers

import (
	"fmt"
	"log"

	"github.com/fullsailor/pkcs7"
)

// ParseCMSSig parses the PKCS7 blob, extracting the certificate common names
func ParseCMSSig(data []byte) {

	p7, err := pkcs7.Parse(data)

	if err != nil {
		log.Printf("Error parsing DER: %s", err.Error())
	}

	if p7 != nil {
		fmt.Printf("CMS Signature has %d certificates:\n", len(p7.Certificates))

		for _, cert := range p7.Certificates {
			fmt.Printf("\tCN: %q\n", cert.Subject.CommonName)
		}
	} else {
		fmt.Printf("No certificates found")
	}
}
