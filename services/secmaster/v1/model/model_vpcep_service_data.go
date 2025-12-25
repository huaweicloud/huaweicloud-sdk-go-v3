package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcepServiceData struct {

	// 已弃用
	Deprecated *bool `json:"deprecated,omitempty"`

	// UUID
	Id *string `json:"id,omitempty"`

	// 所属租户名称
	Name *string `json:"name,omitempty"`

	Type *VpcServiceType `json:"type,omitempty"`
}

func (o VpcepServiceData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcepServiceData struct{}"
	}

	return strings.Join([]string{"VpcepServiceData", string(data)}, " ")
}
