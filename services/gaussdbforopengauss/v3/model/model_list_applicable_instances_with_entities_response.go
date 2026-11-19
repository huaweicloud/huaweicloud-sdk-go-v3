package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicableInstancesWithEntitiesResponse Response Object
type ListApplicableInstancesWithEntitiesResponse struct {

	// **参数解释**: 可应用的实例列表。
	Instances *[]InstancesListWithEntitiesResult `json:"instances,omitempty"`

	// **参数解释**: 查询数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListApplicableInstancesWithEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicableInstancesWithEntitiesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicableInstancesWithEntitiesResponse", string(data)}, " ")
}
