package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaRebalanceLogResponse Response Object
type ShowKafkaRebalanceLogResponse struct {

	// 日志ID。
	Id *string `json:"id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instanceId,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 日志流ID。
	LogStreamId *string `json:"logStreamId,omitempty"`

	// 日志组ID。
	LogGroupId *string `json:"logGroupId,omitempty"`

	// 看板ID。
	DashboardId *string `json:"dashboardId,omitempty"`

	// 创建时间。
	CreateAt *string `json:"createAt,omitempty"`

	// 更新时间。
	UpdateAt       *string `json:"updateAt,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKafkaRebalanceLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaRebalanceLogResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaRebalanceLogResponse", string(data)}, " ")
}
