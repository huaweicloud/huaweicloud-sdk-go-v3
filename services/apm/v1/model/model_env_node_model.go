package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvNodeModel 环境信息。
type EnvNodeModel struct {

	// 环境id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 组件id。
	AppId *int64 `json:"app_id,omitempty"`

	// 应用名称。
	BusinessName *string `json:"business_name,omitempty"`

	// 租户id。
	InnerDomainId *int32 `json:"inner_domain_id,omitempty"`

	// 环境名称。
	Name *string `json:"name,omitempty"`

	// 是否是默认环境。
	IsDefault *bool `json:"is_default,omitempty"`

	// 组件名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 区域。
	Region *string `json:"region,omitempty"`
}

func (o EnvNodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvNodeModel struct{}"
	}

	return strings.Join([]string{"EnvNodeModel", string(data)}, " ")
}
