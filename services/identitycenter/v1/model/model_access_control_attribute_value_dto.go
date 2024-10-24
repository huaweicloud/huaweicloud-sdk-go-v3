package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessControlAttributeValueDto The value used for mapping a specified attribute to an identity source.
type AccessControlAttributeValueDto struct {

	// 用于将指定属性映射到身份源的值
	Source []string `json:"source"`
}

func (o AccessControlAttributeValueDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControlAttributeValueDto struct{}"
	}

	return strings.Join([]string{"AccessControlAttributeValueDto", string(data)}, " ")
}
