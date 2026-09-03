package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupportKeyStringResponse Response Object
type ShowSupportKeyStringResponse struct {

	// 实例是否使用关键字自治限流功能。true：可用，false：不可用
	SupportKeyStr *bool `json:"support_key_str,omitempty"`

	// 实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例详细版本号
	InstanceDetailVersion *string `json:"instance_detail_version,omitempty"`

	// 当support_key_str为False时展示errorMsg
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSupportKeyStringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupportKeyStringResponse struct{}"
	}

	return strings.Join([]string{"ShowSupportKeyStringResponse", string(data)}, " ")
}
