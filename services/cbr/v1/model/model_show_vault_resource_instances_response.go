package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowVaultResourceInstancesResponse struct {
	// 符合查询条件的资源列表（action为count时无此参数）。

	Resources *[]TagResource `json:"resources,omitempty"`
	// 符合查询条件的资源总个数

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultResourceInstancesResponse", string(data)}, " ")
}
