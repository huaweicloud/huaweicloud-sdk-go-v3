package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookInstancesResponse Response Object
type ListPlaybookInstancesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 剧本实例列表信息
	Instances *[]PlaybookInstanceInfo `json:"instances,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookInstancesResponse", string(data)}, " ")
}
