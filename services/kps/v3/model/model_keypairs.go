package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Keypairs struct {
	Keypair *Keypair `json:"keypair" xml:"keypair"`
}

func (o Keypairs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Keypairs struct{}"
	}

	return strings.Join([]string{"Keypairs", string(data)}, " ")
}
