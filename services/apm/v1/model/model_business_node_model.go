package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 业务数据结构
type BusinessNodeModel struct {

	// 默认业务
	Default *bool `json:"default,omitempty" xml:"default"`

	// 业务展示名称
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 企业项目的id
	EpsId *string `json:"eps_id,omitempty" xml:"eps_id"`

	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create"`

	// 修改时间
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify"`

	// 业务id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 内部租户id
	InnerDomainId *int32 `json:"inner_domain_id,omitempty" xml:"inner_domain_id"`

	// 是否是默认的业务
	IsDefault *bool `json:"is_default,omitempty" xml:"is_default"`

	// 业务的英文名称
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o BusinessNodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessNodeModel struct{}"
	}

	return strings.Join([]string{"BusinessNodeModel", string(data)}, " ")
}
