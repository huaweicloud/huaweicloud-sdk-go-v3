package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSyncCommandRequestBody 执行shell命令请求体。
type RunSyncCommandRequestBody struct {

	// ADB命令，固定填写shell。
	Command string `json:"command"`

	// 待执行的命令。  最大长度为1024字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）。
	Content string `json:"content"`

	// 云手机ID列表。 server_ids参数不存在时必选，同时存在只处理phone_ids。最多支持传入15个phone_id。
	PhoneIds *[]string `json:"phone_ids,omitempty"`

	// 云手机服务器ID列表。 phone_ids参数不存在时必选，同时存在只处理phone_ids。
	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o RunSyncCommandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSyncCommandRequestBody struct{}"
	}

	return strings.Join([]string{"RunSyncCommandRequestBody", string(data)}, " ")
}
