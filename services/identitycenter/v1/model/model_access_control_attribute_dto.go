package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessControlAttributeDto These are IAM Identity Center identity store attributes that you can configure for use in attributes-based access control (ABAC).
type AccessControlAttributeDto struct {

	// 与您的身份源中的身份关联的属性的名称
	Key string `json:"key"`

	Value *AccessControlAttributeValueDto `json:"value"`
}

func (o AccessControlAttributeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControlAttributeDto struct{}"
	}

	return strings.Join([]string{"AccessControlAttributeDto", string(data)}, " ")
}
