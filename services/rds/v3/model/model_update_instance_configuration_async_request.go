package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceConfigurationAsyncRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateInstanceConfigurationRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateInstanceConfigurationAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationAsyncRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationAsyncRequest", string(data)}, " ")
}
