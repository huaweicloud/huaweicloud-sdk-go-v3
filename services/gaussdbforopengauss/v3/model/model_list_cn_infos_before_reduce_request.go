package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCnInfosBeforeReduceRequest Request Object
type ListCnInfosBeforeReduceRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListCnInfosBeforeReduceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCnInfosBeforeReduceRequest struct{}"
	}

	return strings.Join([]string{"ListCnInfosBeforeReduceRequest", string(data)}, " ")
}
