package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleInstancesResponse Response Object
type ShowRecycleInstancesResponse struct {

	// **参数解释**： 保留天数。 **取值范围**： 1~7。
	RetentionDays *int32 `json:"retention_days,omitempty"`

	// **参数解释**： 是否使用回收站。 **取值范围**： - true：使用回收站。 - false：不使用回收站。
	DefaultUseRecycle *bool `json:"default_use_recycle,omitempty"`

	// **参数解释**： 回收实例列表。
	RecycleInstances *[]InstanceRecycleInfo `json:"recycle_instances,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ShowRecycleInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowRecycleInstancesResponse", string(data)}, " ")
}
