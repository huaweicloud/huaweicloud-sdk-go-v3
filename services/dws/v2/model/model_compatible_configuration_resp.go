package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompatibleConfigurationResp struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 连接
	Links *[]LinkResp `json:"links,omitempty"`
}

func (o CompatibleConfigurationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleConfigurationResp struct{}"
	}

	return strings.Join([]string{"CompatibleConfigurationResp", string(data)}, " ")
}
