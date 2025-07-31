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

	// TTSS支持配置的词表类型 * CHINESE_G2P:拼音 * PHONETIC_SYMBOL:音标 * CONTINUUM:连读 * ALIAS:别名 * SAY_AS:数字英文读法
	Type *string `json:"type,omitempty"`

	// 分组id
	GroupId *string `json:"group_id,omitempty"`

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
