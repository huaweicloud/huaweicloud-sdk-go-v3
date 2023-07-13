package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSignatureRequest Request Object
type ShowSignatureRequest struct {

	// 签名ID
	Id string `json:"id"`
}

func (o ShowSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSignatureRequest struct{}"
	}

	return strings.Join([]string{"ShowSignatureRequest", string(data)}, " ")
}
