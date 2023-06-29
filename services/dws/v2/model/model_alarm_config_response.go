package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmConfigResponse 告警配置详情
type AlarmConfigResponse struct {

	// 告警配置ID
	Id *string `json:"id,omitempty"`

	// 告警ID
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称
	AlarmName *string `json:"alarm_name,omitempty"`

	// 所属服务，支持DWS,DLI,DGC,CloudTable,CDM,GES,CSS
	NameSpace *string `json:"name_space,omitempty"`

	// 告警级别
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 用户是否可见
	IsUserVisible *string `json:"is_user_visible,omitempty"`

	// 是否覆盖
	IsConverge *string `json:"is_converge,omitempty"`

	// 覆盖时间
	ConvergeTime *int32 `json:"converge_time,omitempty"`

	// 运维是否可见
	IsMaintainVisible *string `json:"is_maintain_visible,omitempty"`
}

func (o AlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"AlarmConfigResponse", string(data)}, " ")
}
