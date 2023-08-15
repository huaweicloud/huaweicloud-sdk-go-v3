package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesFileAndClipboardFileRedirection 文件重定向。
type PoliciesFileAndClipboardFileRedirection struct {

	// 文件重定向。取值为： DISABLED：表示禁用。（默认） READ_ONLY：表示只读。 READ_AND_WRITE：表示读写。
	RedirectionMode *PoliciesFileAndClipboardFileRedirectionRedirectionMode `json:"redirection_mode,omitempty"`

	Options *PoliciesFileAndClipboardFileRedirectionOptions `json:"options,omitempty"`

	// 是否开启发送文件（虚机到客户端）。取值为： false：表示关闭。 true：表示开启。
	VmSendFileClient *bool `json:"vm_send_file_client,omitempty"`

	RedirectionSendFileOptions *PoliciesFileAndClipboardFileRedirectionRedirectionSendFileOptions `json:"redirection_send_file_options,omitempty"`
}

func (o PoliciesFileAndClipboardFileRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboardFileRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboardFileRedirection", string(data)}, " ")
}

type PoliciesFileAndClipboardFileRedirectionRedirectionMode struct {
	value string
}

type PoliciesFileAndClipboardFileRedirectionRedirectionModeEnum struct {
	DISABLED       PoliciesFileAndClipboardFileRedirectionRedirectionMode
	READ_ONLY      PoliciesFileAndClipboardFileRedirectionRedirectionMode
	READ_AND_WRITE PoliciesFileAndClipboardFileRedirectionRedirectionMode
}

func GetPoliciesFileAndClipboardFileRedirectionRedirectionModeEnum() PoliciesFileAndClipboardFileRedirectionRedirectionModeEnum {
	return PoliciesFileAndClipboardFileRedirectionRedirectionModeEnum{
		DISABLED: PoliciesFileAndClipboardFileRedirectionRedirectionMode{
			value: "DISABLED",
		},
		READ_ONLY: PoliciesFileAndClipboardFileRedirectionRedirectionMode{
			value: "READ_ONLY",
		},
		READ_AND_WRITE: PoliciesFileAndClipboardFileRedirectionRedirectionMode{
			value: "READ_AND_WRITE",
		},
	}
}

func (c PoliciesFileAndClipboardFileRedirectionRedirectionMode) Value() string {
	return c.value
}

func (c PoliciesFileAndClipboardFileRedirectionRedirectionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesFileAndClipboardFileRedirectionRedirectionMode) UnmarshalJSON(b []byte) error {
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
