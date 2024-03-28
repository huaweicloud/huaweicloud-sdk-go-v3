package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesByTagResponse Response Object
type CountResourcesByTagResponse struct {

	// 资源总数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountResourcesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagResponse struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagResponse", string(data)}, " ")
}
