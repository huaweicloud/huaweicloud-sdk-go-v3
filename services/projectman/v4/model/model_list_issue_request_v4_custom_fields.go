package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListIssueRequestV4CustomFields struct {

	// 自定义属性字段
	CustomField *string `json:"custom_field,omitempty" xml:"custom_field"`

	// 自定义属性对应的值
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ListIssueRequestV4CustomFields) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueRequestV4CustomFields struct{}"
	}

	return strings.Join([]string{"ListIssueRequestV4CustomFields", string(data)}, " ")
}
