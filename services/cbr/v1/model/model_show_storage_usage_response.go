package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStorageUsageResponse struct {

	// 满足过滤条件的资源总条数
	ResourceCount *int32 `json:"resource_count,omitempty"`

	//
	StorageUsage   *[]StorageUsage `json:"storage_usage,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowStorageUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStorageUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowStorageUsageResponse", string(data)}, " ")
}
