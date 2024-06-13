package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTaskSettingRequest Request Object
type SaveTaskSettingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *SaveTaskSettingRequestBody `json:"body,omitempty"`
}

func (o SaveTaskSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTaskSettingRequest struct{}"
	}

	return strings.Join([]string{"SaveTaskSettingRequest", string(data)}, " ")
}
