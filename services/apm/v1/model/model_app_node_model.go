package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppNodeModel 组件信息。
type AppNodeModel struct {

	// 组件id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 组件名称。
	Name *string `json:"name,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 子应用id。
	SubBusinessId *int64 `json:"sub_business_id,omitempty"`

	// 租户id。
	InnerDomainId *int32 `json:"inner_domain_id,omitempty"`
}

func (o AppNodeModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppNodeModel struct{}"
	}

	return strings.Join([]string{"AppNodeModel", string(data)}, " ")
}
