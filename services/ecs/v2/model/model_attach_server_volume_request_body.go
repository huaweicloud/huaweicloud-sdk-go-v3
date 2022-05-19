package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type AttachServerVolumeRequestBody struct {
	VolumeAttachment *AttachServerVolumeOption `json:"volumeAttachment"`

	// 是否只预检此次请求。  true：发送检查请求，不会挂载磁盘。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应结果。 false：发送正常请求，通过检查后并且进行挂载磁盘请求。 默认值：false
	DryRun *bool `json:"dry_run,omitempty"`
}

func (o AttachServerVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeRequestBody", string(data)}, " ")
}
