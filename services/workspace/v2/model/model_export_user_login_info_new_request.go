package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserLoginInfoNewRequest Request Object
type ExportUserLoginInfoNewRequest struct {

	// 查询的起始时间。指定该参数后，返回的结果为此时间之后的所有登录记录。时间格式如：“2016-08-20T21:11Z”。终止时间不为空时，起始时间为必填参数。
	StartTime *string `json:"start_time,omitempty"`

	// 查询的终止时间。指定该参数后，返回的结果为此时间之前的所有登录记录。时间格式如：“2016-08-20T21:11Z”。起始时间不为空时，终止时间为必填参数。
	EndTime *string `json:"end_time,omitempty"`

	// 登录桌面的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 计算机名（操作系统信息中可见）。
	ComputerName *string `json:"computer_name,omitempty"`

	// 登录桌面的终端系统类型。
	TerminalType *string `json:"terminal_type,omitempty"`

	// 导出语言，默认英文。 - zh_CN：中文 - en_US：英文
	Language *string `json:"language,omitempty"`
}

func (o ExportUserLoginInfoNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserLoginInfoNewRequest struct{}"
	}

	return strings.Join([]string{"ExportUserLoginInfoNewRequest", string(data)}, " ")
}
