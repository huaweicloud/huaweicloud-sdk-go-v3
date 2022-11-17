package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceSet struct {

	// 服务组id
	SetId *string `json:"set_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 引用次数
	RefCount *int32 `json:"ref_count,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`
}

func (o ServiceSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSet struct{}"
	}

	return strings.Join([]string{"ServiceSet", string(data)}, " ")
}
