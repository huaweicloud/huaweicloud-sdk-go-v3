package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserConnectionNewRequest Request Object
type ExportUserConnectionNewRequest struct {

	// 查询的起始时间。指定该参数后，返回的结果为此时间之后的所有登录记录。时间格式如：“2016-08-20T21:11:11Z”。终止时间不为空时，起始时间为非必填项。开始时间要在最近30天内。
	StartTime *string `json:"start_time,omitempty"`

	// 查询的终止时间。指定该参数后，返回的结果为此时间之前的所有登录记录。时间格式如：“2016-08-20T21:11:11Z”。起始时间不为空时，起始时间为非必填项。结束时间要在最近30天内。
	EndTime *string `json:"end_time,omitempty"`

	// 登录桌面的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 计算机名（操作系统信息中可见）。
	ComputerName *string `json:"computer_name,omitempty"`

	// 登录桌面的终端系统类型，该字段支持模糊查询，例如：Windows 10。
	TerminalType *string `json:"terminal_type,omitempty"`

	// 导出语言，默认英文。 - zh_CN：中文 - en_US：英文
	Language *string `json:"language,omitempty"`

	// 查询端到端时延的最小值。
	MinE2eRtt *int32 `json:"min_e2e_rtt,omitempty"`

	// 查询端到端时延的最大值。
	MaxE2eRtt *int32 `json:"max_e2e_rtt,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 客户端出口IP。
	PublicIp *string `json:"public_ip,omitempty"`
}

func (o ExportUserConnectionNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserConnectionNewRequest struct{}"
	}

	return strings.Join([]string{"ExportUserConnectionNewRequest", string(data)}, " ")
}
