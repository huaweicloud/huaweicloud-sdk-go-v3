package signer

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer/algorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSigner(t *testing.T) {
	cases := []struct {
		Name      string
		Algorithm algorithm.SigningAlgorithm
		Signer    IAKSKSigner
	}{
		{"HmacSHA256", algorithm.HmacSHA256, Signer{}},
		{"HmacSM3", algorithm.HmacSM3, SM3Signer{}},
		{"EcdsaP256SHA256", algorithm.EcdsaP256SHA256, P256SHA256Signer{}},
		{"SM2SM3", algorithm.SM2SM3, SM2SM3Signer{}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			signer, err := GetSigner(c.Algorithm)
			assert.Nil(t, err)
			assert.Equal(t, c.Signer, signer)
		})
	}
}
