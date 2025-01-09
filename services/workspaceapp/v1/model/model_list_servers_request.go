package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersRequest Request Object
type ListServersRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 服务器组唯一标识。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 服务器名称，支持部分匹配。
	ServerName *string `json:"server_name,omitempty"`

	// 机器名称，支持部分匹配。
	MachineName *string `json:"machine_name,omitempty"`

	// ip地址，支持部分匹配。
	IpAddr *string `json:"ip_addr,omitempty"`

	// 服务器唯一标识。
	ServerId *string `json:"server_id,omitempty"`

	// 服务器维护状态： - true : 维护态的实例。 - false: 非维护态的实例。
	MaintainStatus *string `json:"maintain_status,omitempty"`

	// 是否是弹性创建： true : 通过弹性伸缩创建。 false: 不是通过弹性伸缩创建。
	ScalingAutoCreate *string `json:"scaling_auto_create,omitempty"`
}

func (o ListServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersRequest struct{}"
	}

	return strings.Join([]string{"ListServersRequest", string(data)}, " ")
}
