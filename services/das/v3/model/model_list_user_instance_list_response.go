package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserInstanceListResponse Response Object
type ListUserInstanceListResponse struct {

	// 实例列表
	DasInstances *[]DasUserInstanceInfo `json:"das_instances,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserInstanceListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserInstanceListResponse struct{}"
	}

	return strings.Join([]string{"ListUserInstanceListResponse", string(data)}, " ")
}
