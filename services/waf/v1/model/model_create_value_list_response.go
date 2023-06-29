package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateValueListResponse Response Object
type CreateValueListResponse struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// 引用表类型
	Type *string `json:"type,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`

	// 引用表时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源：  - 1:表示来源于用户手动创建  - 2:表示来源于智能访问控制自动创建
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateValueListResponse struct{}"
	}

	return strings.Join([]string{"CreateValueListResponse", string(data)}, " ")
}
