package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstancesNumByTagsResponse Response Object
type ShowInstancesNumByTagsResponse struct {

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstancesNumByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstancesNumByTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstancesNumByTagsResponse", string(data)}, " ")
}
