package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenClassification struct {

	// 分类名称
	Name string `json:"name"`

	// 分类描述
	Description *string `json:"description,omitempty"`

	// 分类创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 分类创建时间
	CreateTime float32 `json:"create_time,omitempty"`

	// 分类更新时间
	UpdateTime float32 `json:"update_time,omitempty"`

	// 分类更新者
	UpdateUser *string `json:"update_user,omitempty"`

	// 分类的guid标志
	Guid *string `json:"guid,omitempty"`
}

func (o OpenClassification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenClassification struct{}"
	}

	return strings.Join([]string{"OpenClassification", string(data)}, " ")
}
