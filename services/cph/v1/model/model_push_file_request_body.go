package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushFileRequestBody 推送文件请求体。
type PushFileRequestBody struct {

	// 推送文件固定填写push。
	Command string `json:"command"`

	// 推送的文件只支持tar文件类型，指定OBS桶中的tar文件。 最大长度为1024字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）。文件格式：obs://obs-bucket-name/obs-file-path/file.tar，文件路径与OBS桶对象路径对应。
	Content string `json:"content"`

	// 云手机ID列表。 server_ids参数不存在时必选，同时存在只处理phone_ids。
	PhoneIds *[]string `json:"phone_ids,omitempty"`

	// 云手机服务器ID列表。 phone_ids参数不存在时必选，同时存在只处理phone_ids。
	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o PushFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushFileRequestBody struct{}"
	}

	return strings.Join([]string{"PushFileRequestBody", string(data)}, " ")
}
