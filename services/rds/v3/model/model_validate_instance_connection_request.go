package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateInstanceConnectionRequest Request Object
type ValidateInstanceConnectionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ValidateInstanceConnectionRequestBody `json:"body,omitempty"`
}

func (o ValidateInstanceConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateInstanceConnectionRequest struct{}"
	}

	return strings.Join([]string{"ValidateInstanceConnectionRequest", string(data)}, " ")
}
