package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppliedInstancesResponse Response Object
type ListAppliedInstancesResponse struct {

	// **参数解释：** 可以应用的实例列表。 **取值范围：** 不涉及。
	Instances      *[]ApplicableInstancesInfo `json:"instances,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListAppliedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppliedInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAppliedInstancesResponse", string(data)}, " ")
}
