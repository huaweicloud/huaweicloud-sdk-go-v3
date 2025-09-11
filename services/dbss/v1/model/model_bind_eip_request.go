package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindEipRequest struct {

	// EIP信息
	Eip string `json:"eip"`

	// EIP ID
	EipId string `json:"eip_id"`
}

func (o BindEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindEipRequest struct{}"
	}

	return strings.Join([]string{"BindEipRequest", string(data)}, " ")
}
