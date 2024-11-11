package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("DOMAIN:_, HasMX:_, HasSPF:_, SPRRecord:_, HasDMARC:_, DMARCRecord:_")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Could not read from input: %v", err)
	}
}

func checkDomain(domain string) {
	var HasMx, HasSPF, HasDMARC bool
	var SPFRecord, DMARCRecord string

	nxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(nxRecords) > 0 {
		HasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, v := range txtRecords {
		if strings.HasPrefix(v, "v=spf1") {
			HasSPF = true
			SPFRecord = v
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, v := range dmarcRecords {
		if strings.HasPrefix(v, "v=DMARC1") {
			HasDMARC = true
			DMARCRecord = v
			break
		}
	}

	fmt.Printf("\nDOMAIN: %v\n\nHasMX: %v\n\nHasSPF: %v\n\nSPFRecord: %v\n\nHasDMARC: %v\n\nDMARCRecord: %v\n\n", domain, HasMx, HasSPF, SPFRecord, HasDMARC, DMARCRecord)
}
