package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFreeResourcesUsageRecordsResponse Response Object
type ListFreeResourcesUsageRecordsResponse struct {

	// 资源包使用明细记录，具体参见表1。
	FreeResourceRecords *[]FreeResourceRecord `json:"free_resource_records,omitempty"`

	// 满足条件的总个数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFreeResourcesUsageRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourcesUsageRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListFreeResourcesUsageRecordsResponse", string(data)}, " ")
}
