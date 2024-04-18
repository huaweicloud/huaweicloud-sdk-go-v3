package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceTagRequest Request Object
type DeleteInstanceTagRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 标签键
	Key string `json:"key"`
}

func (o DeleteInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceTagRequest", string(data)}, " ")
}
