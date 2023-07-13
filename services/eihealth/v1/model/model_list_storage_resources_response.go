package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageResourcesResponse Response Object
type ListStorageResourcesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 存储资源列表
	Resources      *[]StorageResourceRsp `json:"resources,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListStorageResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListStorageResourcesResponse", string(data)}, " ")
}
