package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesFileAndClipboardClipboardRedirectionOptions 剪切板重定向控制的选项。
type PoliciesFileAndClipboardClipboardRedirectionOptions struct {

	// 是否开启剪切板富文本重定向。取值为： false：表示关闭。 true：表示开启。
	RichTextRedirectionEnable *bool `json:"rich_text_redirection_enable,omitempty"`

	// 剪切板富文本重定向。取值为： DISABLED：表示禁用。（默认） SERVER_TO_CLIENT_ENABLED：表示开启服务端到客户端。 CLIENT_TO_SERVER_ENABLED：表示开启客户端到服务端。 TWO_WAY_ENABLED：表示开启双向。
	RichTextClipboardRedirection *PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection `json:"rich_text_clipboard_redirection,omitempty"`

	// 是否开启剪切板文件重定向。取值为： false：表示关闭。 true：表示开启。
	ClipboardFileRedirectionEnable *bool `json:"clipboard_file_redirection_enable,omitempty"`

	// 剪切板文件重定向。取值为： DISABLED：表示禁用。（默认） SERVER_TO_CLIENT_ENABLED：表示开启服务端到客户端。 CLIENT_TO_SERVER_ENABLED：表示开启客户端到服务端。 TWO_WAY_ENABLED：表示开启双向。
	FileClipboardRedirection *PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection `json:"file_clipboard_redirection,omitempty"`

	// 本地到虚拟机长度限制开关。取值为： false: 标识关闭。 ture: 标识开启。
	ClipboardLengthLimitCtsEnable *bool `json:"clipboard_length_limit_cts_enable,omitempty"`

	// 本地到虚拟机长度限制。
	ClipboardLengthLimitCts *int32 `json:"clipboard_length_limit_cts,omitempty"`

	// 虚拟机到本地到长度限制开关。取值为： false: 标识关闭。 ture: 标识开启。
	ClipboardLengthLimitStcEnable *bool `json:"clipboard_length_limit_stc_enable,omitempty"`

	// 虚拟机到本地长度限制。
	ClipboardLengthLimitStc *int32 `json:"clipboard_length_limit_stc,omitempty"`
}

func (o PoliciesFileAndClipboardClipboardRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboardClipboardRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboardClipboardRedirectionOptions", string(data)}, " ")
}

type PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection struct {
	value string
}

type PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirectionEnum struct {
	DISABLED                 PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection
	SERVER_TO_CLIENT_ENABLED PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection
	CLIENT_TO_SERVER_ENABLED PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection
	TWO_WAY_ENABLED          PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection
}

func GetPoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirectionEnum() PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirectionEnum {
	return PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirectionEnum{
		DISABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection{
			value: "DISABLED",
		},
		SERVER_TO_CLIENT_ENABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection{
			value: "SERVER_TO_CLIENT_ENABLED",
		},
		CLIENT_TO_SERVER_ENABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection{
			value: "CLIENT_TO_SERVER_ENABLED",
		},
		TWO_WAY_ENABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection{
			value: "TWO_WAY_ENABLED",
		},
	}
}

func (c PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection) Value() string {
	return c.value
}

func (c PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesFileAndClipboardClipboardRedirectionOptionsRichTextClipboardRedirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection struct {
	value string
}

type PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirectionEnum struct {
	DISABLED                 PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection
	SERVER_TO_CLIENT_ENABLED PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection
	CLIENT_TO_SERVER_ENABLED PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection
	TWO_WAY_ENABLED          PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection
}

func GetPoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirectionEnum() PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirectionEnum {
	return PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirectionEnum{
		DISABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection{
			value: "DISABLED",
		},
		SERVER_TO_CLIENT_ENABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection{
			value: "SERVER_TO_CLIENT_ENABLED",
		},
		CLIENT_TO_SERVER_ENABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection{
			value: "CLIENT_TO_SERVER_ENABLED",
		},
		TWO_WAY_ENABLED: PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection{
			value: "TWO_WAY_ENABLED",
		},
	}
}

func (c PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection) Value() string {
	return c.value
}

func (c PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesFileAndClipboardClipboardRedirectionOptionsFileClipboardRedirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
