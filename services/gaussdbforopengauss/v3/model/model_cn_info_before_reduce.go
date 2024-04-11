package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CnInfoBeforeReduce struct {

	// 节点ID。
	Id *string `json:"id,omitempty"`

	// 节点名称。
	Name *string `json:"name,omitempty"`

	// 节点状态。 取值： 值为“normal”，表示节点正常。 值为“abnormal”，表示节点异常。 值为“creating”，表示节点正在创建。 值为“createfail”，表示节点创建失败。
	Status *string `json:"status,omitempty"`

	// 可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 是否允许删除。
	SupportReduce *bool `json:"support_reduce,omitempty"`
}

func (o CnInfoBeforeReduce) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CnInfoBeforeReduce struct{}"
	}

	return strings.Join([]string{"CnInfoBeforeReduce", string(data)}, " ")
}
