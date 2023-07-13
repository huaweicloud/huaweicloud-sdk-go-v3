package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeypairRequest Request Object
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
