package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueueResourcesResponse Response Object
type UpdateQueueResourcesResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateQueueResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueueResourcesResponse struct{}"
	}

	return strings.Join([]string{"UpdateQueueResourcesResponse", string(data)}, " ")
}
