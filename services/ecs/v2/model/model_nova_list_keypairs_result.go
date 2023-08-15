package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListKeypairsResult
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
