package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组件信息
type AppNodeModel struct {

	// 组件id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create"`

	// 修改时间
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify"`

	// 组件名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用id
	BusinessId *int64 `json:"business_id,omitempty" xml:"business_id"`

	// 子应用id
	SubBusinessId *int64 `json:"sub_business_id,omitempty" xml:"sub_business_id"`

	// 租户id
	InnerDomainId *int32 `json:"inner_domain_id,omitempty" xml:"inner_domain_id"`
}

func (o AppNodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppNodeModel struct{}"
	}

	return strings.Join([]string{"AppNodeModel", string(data)}, " ")
}
