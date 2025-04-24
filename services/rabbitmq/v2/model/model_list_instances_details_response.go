package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesDetailsResponse Response Object
type ListInstancesDetailsResponse struct {

	// **参数解释**： 实例列表。
	Instances *[]ShowInstanceResp `json:"instances,omitempty"`

	// **参数解释**： 实例个数。 **取值范围**： 不涉及。
	InstanceNum    *int32 `json:"instance_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesDetailsResponse", string(data)}, " ")
}
