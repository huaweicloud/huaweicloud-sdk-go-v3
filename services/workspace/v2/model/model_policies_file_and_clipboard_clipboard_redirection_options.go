package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesFileAndClipboardClipboardRedirectionOptions 剪切板重定向控制的选项。
type PoliciesFileAndClipboardClipboardRedirectionOptions struct {

	// 是否开启剪切板富文本重定向。取值为： false：表示关闭。 true：表示开启。
	RichTextRedirectionEnable *bool `json:"rich_text_redirection_enable,omitempty"`

	// 是否开启剪切板文件重定向。取值为： false：表示关闭。 true：表示开启。
	ClipboardFileRedirectionEnable *bool `json:"clipboard_file_redirection_enable,omitempty"`
}

func (o PoliciesFileAndClipboardClipboardRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboardClipboardRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboardClipboardRedirectionOptions", string(data)}, " ")
}
