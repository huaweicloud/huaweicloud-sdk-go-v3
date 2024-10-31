package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpReferenceTableResponse Response Object
type UpdateHttpReferenceTableResponse struct {

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

	// 创建来源
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateHttpReferenceTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpReferenceTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpReferenceTableResponse", string(data)}, " ")
}
