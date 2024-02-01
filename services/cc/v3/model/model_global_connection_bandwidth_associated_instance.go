package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlobalConnectionBandwidthAssociatedInstance struct {

	// 绑定实例ID。
	Id string `json:"id"`

	// 绑定实例类型。
	Type string `json:"type"`

	// 绑定实例的region信息，global服务默认取值\"global\"。
	RegionId *string `json:"region_id,omitempty"`

	// 绑定实例的project id信息。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o GlobalConnectionBandwidthAssociatedInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthAssociatedInstance struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthAssociatedInstance", string(data)}, " ")
}
