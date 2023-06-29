package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceResponse Response Object
type ListInstanceResponse struct {

	// 实例总数
	Count *int32 `json:"count,omitempty"`

	// 实例详情列表
	Databases      *[]DatabaseDto `json:"databases,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceResponse", string(data)}, " ")
}
