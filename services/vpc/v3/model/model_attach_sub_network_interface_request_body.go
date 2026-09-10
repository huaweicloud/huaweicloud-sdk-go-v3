package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachSubNetworkInterfaceRequestBody struct {

	// **参数解释**： 是否只预检此次请求。 **约束限制**： 不涉及。 **取值范围**： - true：发送检查请求，不会挂载辅助弹性网卡。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。 - false：发送正常请求，并直接挂载辅助弹性网卡。 **默认取值**： false
	DryRun *bool `json:"dry_run,omitempty"`

	SubNetworkInterface *AttachSubNetworkInterfaceOption `json:"sub_network_interface,omitempty"`
}

func (o AttachSubNetworkInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachSubNetworkInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"AttachSubNetworkInterfaceRequestBody", string(data)}, " ")
}
