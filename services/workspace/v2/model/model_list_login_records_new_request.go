package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginRecordsNewRequest Request Object
type ListLoginRecordsNewRequest struct {

	// 查询的起始时间。指定该参数后，返回的结果为此时间之后的所有登录记录。时间格式如：“2016-08-20T21:11Z”。终止时间不为空时，起始时间为必填参数。
	StartTime *string `json:"start_time,omitempty"`

	// 查询的终止时间。指定该参数后，返回的结果为此时间之前的所有登录记录。时间格式如：“2016-08-20T21:11Z”。起始时间不为空时，终止时间为必填参数。
	EndTime *string `json:"end_time,omitempty"`

	// 登录桌面的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 计算机名（操作系统信息中可见）。
	ComputerName *string `json:"computer_name,omitempty"`

	// 登录桌面的终端系统类型，当前支持：WI（云桌面客户端）。
	TerminalType *string `json:"terminal_type,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 终端出口IP。
	PublicIp *string `json:"public_ip,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回登录记录数量限制,取值范围0-1000。如果不指定，默认为20。
	Limit *int32 `json:"limit,omitempty"`

	// 查询端到端时延的最小值。
	MinNetworkRtt *int32 `json:"min_network_rtt,omitempty"`

	// 查询端到端时延的最大值。
	MaxNetworkRtt *int32 `json:"max_network_rtt,omitempty"`
}

func (o ListLoginRecordsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginRecordsNewRequest struct{}"
	}

	return strings.Join([]string{"ListLoginRecordsNewRequest", string(data)}, " ")
}
