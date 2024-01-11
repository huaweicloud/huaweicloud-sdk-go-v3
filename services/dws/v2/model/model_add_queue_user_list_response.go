package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddQueueUserListResponse Response Object
type AddQueueUserListResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddQueueUserListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddQueueUserListResponse struct{}"
	}

	return strings.Join([]string{"AddQueueUserListResponse", string(data)}, " ")
}
