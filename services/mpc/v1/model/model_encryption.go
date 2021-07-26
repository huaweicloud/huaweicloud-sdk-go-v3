package model

import (
	"encoding/json"

	"strings"
)

type Encryption struct {
	HlsEncrypt *HlsEncrypt `json:"hls_encrypt,omitempty"`
}

func (o Encryption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Encryption struct{}"
	}

	return strings.Join([]string{"Encryption", string(data)}, " ")
}
