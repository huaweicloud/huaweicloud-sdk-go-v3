package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesFileAndClipboardFileRedirectionRedirectionSendFileOptions 重定向和发送文件共同控制项。
type PoliciesFileAndClipboardFileRedirectionRedirectionSendFileOptions struct {

	// 读写速度（Kbps）。取值范围为[0-2147483647]。默认：0。
	ReadWriteSpeed *int32 `json:"read_write_speed,omitempty"`
}

func (o PoliciesFileAndClipboardFileRedirectionRedirectionSendFileOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboardFileRedirectionRedirectionSendFileOptions struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboardFileRedirectionRedirectionSendFileOptions", string(data)}, " ")
}
