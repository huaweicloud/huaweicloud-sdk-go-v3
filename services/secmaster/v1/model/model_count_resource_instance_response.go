package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourceInstanceResponse Response Object
type CountResourceInstanceResponse struct {

	// 资产总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountResourceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstanceResponse struct{}"
	}

	return strings.Join([]string{"CountResourceInstanceResponse", string(data)}, " ")
}
