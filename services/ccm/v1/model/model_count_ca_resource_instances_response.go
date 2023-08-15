package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountCaResourceInstancesResponse Response Object
type CountCaResourceInstancesResponse struct {

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountCaResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountCaResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"CountCaResourceInstancesResponse", string(data)}, " ")
}
