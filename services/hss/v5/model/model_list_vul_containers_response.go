package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulContainersResponse Response Object
type ListVulContainersResponse struct {

	// 受影响的容器总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 受影响的容器列表
	DataList       *[]VulContainerInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVulContainersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulContainersResponse struct{}"
	}

	return strings.Join([]string{"ListVulContainersResponse", string(data)}, " ")
}
