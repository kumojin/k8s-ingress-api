package network

import (
	"net"
)

func ValidateCNAME(host, expectedCNAME string) (bool, error) {
	result, err := net.LookupCNAME(host)
	if err != nil {
		return false, err
	}
	return result[:len(result)-1] == expectedCNAME, nil
}
