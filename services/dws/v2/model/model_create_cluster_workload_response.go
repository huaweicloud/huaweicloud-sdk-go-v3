package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterWorkloadResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterWorkloadResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterWorkloadResponse", string(data)}, " ")
}
