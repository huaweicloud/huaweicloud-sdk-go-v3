package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncRequirementsResponse Response Object
type ListSyncRequirementsResponse struct {

	// 表示当前集群是否已完成IAM同步
	IsSynchronous  *bool `json:"is_synchronous,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListSyncRequirementsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncRequirementsResponse struct{}"
	}

	return strings.Join([]string{"ListSyncRequirementsResponse", string(data)}, " ")
}
