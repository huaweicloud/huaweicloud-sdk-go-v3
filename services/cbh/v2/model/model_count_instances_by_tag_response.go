package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInstancesByTagResponse Response Object
type CountInstancesByTagResponse struct {

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountInstancesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInstancesByTagResponse struct{}"
	}

	return strings.Join([]string{"CountInstancesByTagResponse", string(data)}, " ")
}
