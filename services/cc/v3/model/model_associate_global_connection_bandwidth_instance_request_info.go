package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateGlobalConnectionBandwidthInstanceRequestInfo struct {

	// 功能说明：实例ID。 取值范围：1-36个字符，支持数字、字母、_(下划线)、-（中划线）
	ResourceId string `json:"resource_id"`

	// 功能说明：实例类型。
	ResourceType string `json:"resource_type"`

	// 功能说明：实例所在region，不填默认\"global\"。
	RegionId *string `json:"region_id,omitempty"`

	// 功能说明：实例所在region对应的projectId。
	ProjectId string `json:"project_id"`
}

func (o AssociateGlobalConnectionBandwidthInstanceRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGlobalConnectionBandwidthInstanceRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociateGlobalConnectionBandwidthInstanceRequestInfo", string(data)}, " ")
}
