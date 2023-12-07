package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStacksResponse Response Object
type ListStacksResponse struct {

	// 资源栈列表。默认按照生成时间降序排序，最新生成的在最前
	Stacks         *[]Stack `json:"stacks,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListStacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStacksResponse struct{}"
	}

	return strings.Join([]string{"ListStacksResponse", string(data)}, " ")
}
