package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactorySearchBaselineInstancesResponse Response Object
type ListFactorySearchBaselineInstancesResponse struct {

	// 基线实例总数。
	Total *int32 `json:"total,omitempty"`

	// 基线实例列表信息。
	BaselineInstances *[]BaselineInstance `json:"baseline_instances,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListFactorySearchBaselineInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactorySearchBaselineInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListFactorySearchBaselineInstancesResponse", string(data)}, " ")
}
