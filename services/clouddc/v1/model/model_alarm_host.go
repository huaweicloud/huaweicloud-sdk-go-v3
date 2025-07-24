package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmHost struct {

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 告警信息
	AlarmInfos *[]AlarmInfo `json:"alarm_infos,omitempty"`
}

func (o AlarmHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHost struct{}"
	}

	return strings.Join([]string{"AlarmHost", string(data)}, " ")
}
