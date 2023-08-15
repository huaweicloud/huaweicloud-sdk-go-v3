package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageAutoSyncReposDetailsResponse Response Object
type ListImageAutoSyncReposDetailsResponse struct {

	// 镜像自动同步规则
	Body           *[]SyncRepo `json:"body,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListImageAutoSyncReposDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageAutoSyncReposDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListImageAutoSyncReposDetailsResponse", string(data)}, " ")
}
