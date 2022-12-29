package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEastWestFirewallRequest struct {

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListEastWestFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEastWestFirewallRequest struct{}"
	}

	return strings.Join([]string{"ListEastWestFirewallRequest", string(data)}, " ")
}
