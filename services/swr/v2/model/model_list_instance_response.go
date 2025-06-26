package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceResponse Response Object
type ListInstanceResponse struct {

	// 实例列表
	Instances *[]Instance `json:"instances,omitempty"`

	// 实例总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceResponse", string(data)}, " ")
}
