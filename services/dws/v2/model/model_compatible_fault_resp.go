package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompatibleFaultResp struct {

	// 信息
	Message *string `json:"message,omitempty"`

	// 创建者
	Created *string `json:"created,omitempty"`

	// 详细
	Details *string `json:"details,omitempty"`
}

func (o CompatibleFaultResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleFaultResp struct{}"
	}

	return strings.Join([]string{"CompatibleFaultResp", string(data)}, " ")
}
