package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsSchedulerHint 弹性云服务器调度信息。
type HwcEcsSchedulerHint struct {

	// 云服务器组ID，UUID格式。 云服务器组的ID可以从控制台或者参考查询云服务器组列表获取。
	Group *[]string `json:"group,omitempty"`

	// 在指定的专属主机或者共享主机上创建弹性云服务器。 参数值为shared或者dedicated。
	Tenancy *[]string `json:"tenancy,omitempty"`

	// 专属主机的ID。 说明： 专属主机的ID仅在tenancy为dedicated时生效。
	DedicatedHostId *[]string `json:"dedicated_host_id,omitempty"`
}

func (o HwcEcsSchedulerHint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsSchedulerHint struct{}"
	}

	return strings.Join([]string{"HwcEcsSchedulerHint", string(data)}, " ")
}
