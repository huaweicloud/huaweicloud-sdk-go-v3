package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppGroupsResponse Response Object
type ListAppGroupsResponse struct {

	// 分组信息列表
	Result *[]AppGroupsEntity `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAppGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListAppGroupsResponse", string(data)}, " ")
}
