package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceAccessoriesResponse Response Object
type ListInstanceAccessoriesResponse struct {

	// 制品附件列表
	Accessories *[]Accessory `json:"accessories,omitempty"`

	// 制品附件总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceAccessoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceAccessoriesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceAccessoriesResponse", string(data)}, " ")
}
