package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatusServiceRequest Request Object
type ShowStatusServiceRequest struct {

	// 集群ID，默认空字符串，默认查询所有集群。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 应用ID，默认空字符串，默认查询所有应用。
	AppId *string `json:"app_id,omitempty"`

	// 集群ID，默认空字符串，默认查询所有集群。
	InstanceId *string `json:"instance_id,omitempty"`

	// 查询范围起始时间，毫秒时间戳，默认为0，默认查询三天前。
	From *int64 `json:"from,omitempty"`

	// 查询范围终止时间，毫秒时间戳，默认为0，默认查询到当前时间。
	To *int64 `json:"to,omitempty"`

	// 数据统计周期，默认0，默认为5分钟。
	Period *int32 `json:"period,omitempty"`

	// 统计类型，默认min，默认用统计周期中数据的最小值。
	Filter *string `json:"filter,omitempty"`

	// 服务类型，默认空字符串，默认查询所有服务类型。
	ServerType *string `json:"server_type,omitempty"`
}

func (o ShowStatusServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusServiceRequest struct{}"
	}

	return strings.Join([]string{"ShowStatusServiceRequest", string(data)}, " ")
}
