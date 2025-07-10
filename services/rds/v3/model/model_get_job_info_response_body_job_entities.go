package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetJobInfoResponseBodyJobEntities 根据不同的任务，显示不同的内容。
type GetJobInfoResponseBodyJobEntities struct {
	Instance *GetJobInfoResponseBodyJobEntitiesInstance `json:"instance,omitempty"`

	// 创建实例，单转主备，创建只读实例，调整实例容量任务时等任务时返回。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	Volume *GetJobInfoResponseBodyJobEntitiesVolume `json:"volume,omitempty"`

	// 绑定/解绑弹性IP，开启远程连接等任务时返回。
	PublicIp *string `json:"public_ip,omitempty"`

	// 主备倒换任务时返回。
	SwitchStrategy *string `json:"switch_strategy,omitempty"`
}

func (o GetJobInfoResponseBodyJobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetJobInfoResponseBodyJobEntities struct{}"
	}

	return strings.Join([]string{"GetJobInfoResponseBodyJobEntities", string(data)}, " ")
}
