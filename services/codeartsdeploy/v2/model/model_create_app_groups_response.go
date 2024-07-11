package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppGroupsResponse Response Object
type CreateAppGroupsResponse struct {

	// 分组id
	Result *string `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppGroupsResponse struct{}"
	}

	return strings.Join([]string{"CreateAppGroupsResponse", string(data)}, " ")
}
