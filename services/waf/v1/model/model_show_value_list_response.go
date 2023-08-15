package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowValueListResponse Response Object
type ShowValueListResponse struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// 引用表类型
	Type *string `json:"type,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源：  - 1:表示来源于用户手动创建  - 2:表示来源于智能访问控制自动创建
	Producer *int32 `json:"producer,omitempty"`

	// 创建规则的时间，格式为13位毫秒时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowValueListResponse struct{}"
	}

	return strings.Join([]string{"ShowValueListResponse", string(data)}, " ")
}
