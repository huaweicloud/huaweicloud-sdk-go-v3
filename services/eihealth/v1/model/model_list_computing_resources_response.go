package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComputingResourcesResponse Response Object
type ListComputingResourcesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 计算资源列表
	Resources      *[]ComputingResourceRsp `json:"resources,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListComputingResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListComputingResourcesResponse", string(data)}, " ")
}
