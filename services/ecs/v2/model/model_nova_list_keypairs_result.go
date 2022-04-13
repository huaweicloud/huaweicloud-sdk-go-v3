package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NovaListKeypairsResult struct {
	Keypair *NovaSimpleKeypair `json:"keypair"`
}

func (o NovaListKeypairsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListKeypairsResult struct{}"
	}

	return strings.Join([]string{"NovaListKeypairsResult", string(data)}, " ")
}
