package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDiskRequest Request Object
type ShowInstanceDiskRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceDiskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDiskRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDiskRequest", string(data)}, " ")
}
