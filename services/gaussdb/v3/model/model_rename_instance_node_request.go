package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameInstanceNodeRequest Request Object
type RenameInstanceNodeRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RenameInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o RenameInstanceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"RenameInstanceNodeRequest", string(data)}, " ")
}
