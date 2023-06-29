package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroup 策略组
type PolicyGroup struct {

	// 策略组的唯一标识
	Id *string `json:"id,omitempty"`

	// 策略组名称
	Name *string `json:"name,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 服务器组描述
	Description *string `json:"description,omitempty"`

	// 应用对象列表。
	Targets *[]Target `json:"targets,omitempty"`

	Policies *Policies `json:"policies,omitempty"`

	// 策略组创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 策略组更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o PolicyGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroup struct{}"
	}

	return strings.Join([]string{"PolicyGroup", string(data)}, " ")
}
