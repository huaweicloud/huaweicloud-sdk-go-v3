package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowKeypairRequest struct {
	// 密钥名称

	KeypairName string `json:"keypair_name"`
}

func (o ShowKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeypairRequest struct{}"
	}

	return strings.Join([]string{"ShowKeypairRequest", string(data)}, " ")
}
