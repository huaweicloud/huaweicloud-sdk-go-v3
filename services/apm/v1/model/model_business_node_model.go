package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 业务数据结构
type BusinessNodeModel struct {
	// 默认业务

	Default *bool `json:"default,omitempty"`
	// 业务展示名称

	DisplayName *string `json:"display_name,omitempty"`
	// 企业项目的id

	EpsId *string `json:"eps_id,omitempty"`
	// 创建时间

	GmtCreate *string `json:"gmt_create,omitempty"`
	// 修改时间

	GmtModify *string `json:"gmt_modify,omitempty"`
	// 业务id

	Id *int64 `json:"id,omitempty"`
	// 内部租户id

	InnerDomainId *int32 `json:"inner_domain_id,omitempty"`
	// 是否是默认的业务

	IsDefault *bool `json:"is_default,omitempty"`
	// 业务的英文名称

	Name *string `json:"name,omitempty"`
}

func (o BusinessNodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessNodeModel struct{}"
	}

	return strings.Join([]string{"BusinessNodeModel", string(data)}, " ")
}
