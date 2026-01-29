package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionGlobalResource 当前region已购资源
type SubscriptionGlobalResource struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 当前region编码，默认为null，即为当前region
	RegionId *string `json:"region_id,omitempty"`

	// 当前demo-regionon已购资源列表
	Resources *[]SubscriptionResourceInfo `json:"resources,omitempty"`
}

func (o SubscriptionGlobalResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionGlobalResource struct{}"
	}

	return strings.Join([]string{"SubscriptionGlobalResource", string(data)}, " ")
}
