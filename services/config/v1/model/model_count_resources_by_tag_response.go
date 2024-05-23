package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesByTagResponse Response Object
type CountResourcesByTagResponse struct {

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountResourcesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagResponse struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagResponse", string(data)}, " ")
}
