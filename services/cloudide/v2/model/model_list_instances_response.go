package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesResponse Response Object
type ListInstancesResponse struct {
	Instances *PageInstancesVo `json:"instances,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
