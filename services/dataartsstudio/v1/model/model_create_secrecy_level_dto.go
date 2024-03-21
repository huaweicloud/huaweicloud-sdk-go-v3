package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSecrecyLevelDto struct {

	// 密级名称。名称应该唯一，只能由英文字母、数字、下划线、汉字构成。
	Name *string `json:"name,omitempty"`

	// 密级描述
	Description *string `json:"description,omitempty"`
}

func (o CreateSecrecyLevelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecrecyLevelDto struct{}"
	}

	return strings.Join([]string{"CreateSecrecyLevelDto", string(data)}, " ")
}
