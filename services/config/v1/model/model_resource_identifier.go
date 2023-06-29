package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceIdentifier struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务类型。
	Provider string `json:"provider"`

	// 资源类型。
	Type string `json:"type"`

	// 源帐号ID。
	SourceAccountId string `json:"source_account_id"`

	// 资源所属区域。
	RegionId string `json:"region_id"`
}

func (o ResourceIdentifier) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceIdentifier struct{}"
	}

	return strings.Join([]string{"ResourceIdentifier", string(data)}, " ")
}
