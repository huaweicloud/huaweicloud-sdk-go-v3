package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VocabularyConfig struct {

	// id
	Id *string `json:"id,omitempty"`

	// 原始词
	Key *string `json:"key,omitempty"`

	// 设置的自定义读法
	Value *string `json:"value,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o VocabularyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VocabularyConfig struct{}"
	}

	return strings.Join([]string{"VocabularyConfig", string(data)}, " ")
}
