package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WdhParam wdh参数
type WdhParam struct {

	// 在指定的桌面专属主机上创建桌面。  - dedicated：桌面专属主机。
	Tenancy *string `json:"tenancy,omitempty"`

	// 桌面专属主机的ID。 指定桌面专属主机的ID则将桌面放置到此桌面专属主机。 未指定桌面专属主机的ID则使用自动放置功能放置到可用的桌面专属主机。 注意： - 仅在tenancy指定为dedicated时此字段生效。 - 若要使用自动放置功能来创建桌面，您需要先开启桌面专属主机的自动放置功能。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`
}

func (o WdhParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WdhParam struct{}"
	}

	return strings.Join([]string{"WdhParam", string(data)}, " ")
}
