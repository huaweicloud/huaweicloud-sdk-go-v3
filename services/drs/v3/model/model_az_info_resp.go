package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AzInfoResp 可用区信息
type AzInfoResp struct {

	// 可用区标识代码
	Code *string `json:"code,omitempty"`

	// 可用区名称
	Name *string `json:"name,omitempty"`

	// 可用区状态
	Status *string `json:"status,omitempty"`
}

func (o AzInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInfoResp struct{}"
	}

	return strings.Join([]string{"AzInfoResp", string(data)}, " ")
}
