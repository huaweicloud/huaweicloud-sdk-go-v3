package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpecCheckitemModel struct {

	// 检查项的uuid
	CheckitemUuid *string `json:"checkitem_uuid,omitempty"`

	// 遵从包创建时间戳
	CreateTime *string `json:"create_time,omitempty"`

	// 遵从包创建的语言环境
	Language *string `json:"language,omitempty"`

	// 遵从包名称
	Name *string `json:"name,omitempty"`

	// 遵从包删除时间戳
	RemoveTime *string `json:"remove_time,omitempty"`

	// 遵从包的uuid
	SpecificationUuid *string `json:"specification_uuid,omitempty"`

	// uuid
	Uuid *string `json:"uuid,omitempty"`
}

func (o SpecCheckitemModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecCheckitemModel struct{}"
	}

	return strings.Join([]string{"SpecCheckitemModel", string(data)}, " ")
}
