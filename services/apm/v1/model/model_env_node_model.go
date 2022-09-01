package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 环境信息
type EnvNodeModel struct {

	// 环境id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create"`

	// 修改时间
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify"`

	// 组件id
	AppId *int64 `json:"app_id,omitempty" xml:"app_id"`

	// 应用名称
	BusinessName *string `json:"business_name,omitempty" xml:"business_name"`

	// 租户id
	InnerDomainId *int32 `json:"inner_domain_id,omitempty" xml:"inner_domain_id"`

	// 环境名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否是默认环境
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default"`

	// 组件名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 应用id
	BusinessId *int64 `json:"business_id,omitempty" xml:"business_id"`

	// 区域
	Region *string `json:"region,omitempty" xml:"region"`
}

func (o EnvNodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvNodeModel struct{}"
	}

	return strings.Join([]string{"EnvNodeModel", string(data)}, " ")
}
