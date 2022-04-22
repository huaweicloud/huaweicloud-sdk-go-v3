package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPodsResponse struct {

	// 应用实例总数
	Count *int32 `json:"count,omitempty"`

	// 应用实例列表
	Pods           *[]PodResp `json:"pods,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListPodsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPodsResponse struct{}"
	}

	return strings.Join([]string{"ListPodsResponse", string(data)}, " ")
}
