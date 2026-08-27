package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportDesktopVersionListRequest Request Object
type ExportDesktopVersionListRequest struct {

	// 桌面agent版本号（精确匹配）。
	AgentVersion string `json:"agent_version"`

	// 桌面操作系统类型。
	OsType string `json:"os_type"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称（支持模糊匹配）。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 用户名（支持模糊匹配）。
	Username *string `json:"username,omitempty"`

	// 语言。  - zh_CN：中文 - en_US：英文
	Language ExportDesktopVersionListRequestLanguage `json:"language"`
}

func (o ExportDesktopVersionListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesktopVersionListRequest struct{}"
	}

	return strings.Join([]string{"ExportDesktopVersionListRequest", string(data)}, " ")
}

type ExportDesktopVersionListRequestLanguage struct {
	value string
}

type ExportDesktopVersionListRequestLanguageEnum struct {
	ZH_CN ExportDesktopVersionListRequestLanguage
	EN_US ExportDesktopVersionListRequestLanguage
}

func GetExportDesktopVersionListRequestLanguageEnum() ExportDesktopVersionListRequestLanguageEnum {
	return ExportDesktopVersionListRequestLanguageEnum{
		ZH_CN: ExportDesktopVersionListRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportDesktopVersionListRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportDesktopVersionListRequestLanguage) Value() string {
	return c.value
}

func (c ExportDesktopVersionListRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportDesktopVersionListRequestLanguage) UnmarshalJSON(b []byte) error {
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
