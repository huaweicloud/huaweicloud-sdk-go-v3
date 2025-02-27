package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceStatusResponse Response Object
type ShowInstanceStatusResponse struct {

	// 云堡垒机实例名称。
	Name *string `json:"name,omitempty"`

	// 堡垒机实例状态。 - POWERING_ON：正在开启 - POWERING_OFF：正在关闭 - DELETE_WAITE：等待删除 - REBOOTING：重启中 - RESIZE：变更中 - UPGRADING：升级中 - FROZEN：冻结 - ACTIVE：运行
	Status *string `json:"status,omitempty"`

	// 云堡垒机实例ID，使用UUID格式表示。
	ServerId       *string `json:"server_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusResponse", string(data)}, " ")
}
