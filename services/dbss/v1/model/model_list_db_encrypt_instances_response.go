package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbEncryptInstancesResponse Response Object
type ListDbEncryptInstancesResponse struct {

	// 实例列表
	Clusters       *[]Cluster `json:"clusters,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListDbEncryptInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbEncryptInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListDbEncryptInstancesResponse", string(data)}, " ")
}
