package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstanceResponse Response Object
type ListResourceInstanceResponse struct {

	// 资产信息
	Resources *[]ResourceInfo `json:"resources,omitempty"`

	// 资产总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceResponse", string(data)}, " ")
}
