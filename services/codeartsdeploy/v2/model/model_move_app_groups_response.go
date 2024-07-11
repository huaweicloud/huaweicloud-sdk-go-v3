package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAppGroupsResponse Response Object
type MoveAppGroupsResponse struct {

	// 分组id
	Result *string `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MoveAppGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppGroupsResponse struct{}"
	}

	return strings.Join([]string{"MoveAppGroupsResponse", string(data)}, " ")
}
