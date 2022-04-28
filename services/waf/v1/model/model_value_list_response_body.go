package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 引用表
type ValueListResponseBody struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// 引用表类型
	Type *string `json:"type,omitempty"`

	// 引用表时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源，1代表用户创建，其它值代表modulleX自动生成
	Producer *int32 `json:"producer,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`
}

func (o ValueListResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueListResponseBody struct{}"
	}

	return strings.Join([]string{"ValueListResponseBody", string(data)}, " ")
}
