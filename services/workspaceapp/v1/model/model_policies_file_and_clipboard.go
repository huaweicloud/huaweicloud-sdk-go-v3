package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesFileAndClipboard 文件和剪切板。
type PoliciesFileAndClipboard struct {

	// 应用聚合场景下是否双向放通：取值为： false：表示不放通。 true：表示放通。
	BypassInRemoteAppEnable *bool `json:"bypass_in_remote_app_enable,omitempty"`

	FileRedirection *PoliciesFileAndClipboardFileRedirection `json:"file_redirection,omitempty"`

	// 移动客户端文件重定向：取值为： false：表示关闭。 true：表示开启。
	FdMobileClientRedirEnable *bool `json:"fd_mobile_client_redir_enable,omitempty"`

	// 剪切板重定向。取值为： DISABLED：表示禁用。（默认） SERVER_TO_CLIENT_ENABLED：表示开启服务端到客户端。 CLIENT_TO_SERVER_ENABLED：表示开启客户端到服务端。 TWO_WAY_ENABLED：表示开启双向。
	ClipboardRedirection *PoliciesFileAndClipboardClipboardRedirection `json:"clipboard_redirection,omitempty"`

	ClipboardRedirectionOptions *PoliciesFileAndClipboardClipboardRedirectionOptions `json:"clipboard_redirection_options,omitempty"`
}

func (o PoliciesFileAndClipboard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboard struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboard", string(data)}, " ")
}

type PoliciesFileAndClipboardClipboardRedirection struct {
	value string
}

type PoliciesFileAndClipboardClipboardRedirectionEnum struct {
	DISABLED                 PoliciesFileAndClipboardClipboardRedirection
	SERVER_TO_CLIENT_ENABLED PoliciesFileAndClipboardClipboardRedirection
	CLIENT_TO_SERVER_ENABLED PoliciesFileAndClipboardClipboardRedirection
	TWO_WAY_ENABLED          PoliciesFileAndClipboardClipboardRedirection
}

func GetPoliciesFileAndClipboardClipboardRedirectionEnum() PoliciesFileAndClipboardClipboardRedirectionEnum {
	return PoliciesFileAndClipboardClipboardRedirectionEnum{
		DISABLED: PoliciesFileAndClipboardClipboardRedirection{
			value: "DISABLED",
		},
		SERVER_TO_CLIENT_ENABLED: PoliciesFileAndClipboardClipboardRedirection{
			value: "SERVER_TO_CLIENT_ENABLED",
		},
		CLIENT_TO_SERVER_ENABLED: PoliciesFileAndClipboardClipboardRedirection{
			value: "CLIENT_TO_SERVER_ENABLED",
		},
		TWO_WAY_ENABLED: PoliciesFileAndClipboardClipboardRedirection{
			value: "TWO_WAY_ENABLED",
		},
	}
}

func (c PoliciesFileAndClipboardClipboardRedirection) Value() string {
	return c.value
}

func (c PoliciesFileAndClipboardClipboardRedirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesFileAndClipboardClipboardRedirection) UnmarshalJSON(b []byte) error {
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
