package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountGcbResourceByTagResponse Response Object
type CountGcbResourceByTagResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountGcbResourceByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGcbResourceByTagResponse struct{}"
	}

	return strings.Join([]string{"CountGcbResourceByTagResponse", string(data)}, " ")
}
