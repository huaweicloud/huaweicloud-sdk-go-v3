package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResourceInstancesCountResponse struct {

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResourceInstancesCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceInstancesCountResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceInstancesCountResponse", string(data)}, " ")
}
