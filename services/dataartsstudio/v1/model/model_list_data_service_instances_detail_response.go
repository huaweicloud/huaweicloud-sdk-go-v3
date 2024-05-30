package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataServiceInstancesDetailResponse Response Object
type ListDataServiceInstancesDetailResponse struct {

	// 集群数量。
	Total *int32 `json:"total,omitempty"`

	// 集群概览信息。
	Instances      *[]InstanceDetailDto `json:"instances,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListDataServiceInstancesDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceInstancesDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDataServiceInstancesDetailResponse", string(data)}, " ")
}
