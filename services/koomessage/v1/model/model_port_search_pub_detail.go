package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PortSearchPubDetail struct {

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// 服务号备注。
	PubReference *string `json:"pub_reference,omitempty"`
}

func (o PortSearchPubDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortSearchPubDetail struct{}"
	}

	return strings.Join([]string{"PortSearchPubDetail", string(data)}, " ")
}
