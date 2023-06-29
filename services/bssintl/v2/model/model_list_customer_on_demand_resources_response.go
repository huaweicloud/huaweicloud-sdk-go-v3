package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerOnDemandResourcesResponse Response Object
type ListCustomerOnDemandResourcesResponse struct {

	// 客户资源列表。 具体参见表2。
	Resources *[]CustomerOnDemandResource `json:"resources,omitempty"`

	// 查询总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCustomerOnDemandResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerOnDemandResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerOnDemandResourcesResponse", string(data)}, " ")
}
