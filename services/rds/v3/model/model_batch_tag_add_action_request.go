package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchTagAddActionRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *BatchTagActionAddRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchTagAddActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagAddActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagAddActionRequest", string(data)}, " ")
}
