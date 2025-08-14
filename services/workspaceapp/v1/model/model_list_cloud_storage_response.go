package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudStorageResponse Response Object
type ListCloudStorageResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 存储列表。
	Items          *[]CloudStorage `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCloudStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudStorageResponse struct{}"
	}

	return strings.Join([]string{"ListCloudStorageResponse", string(data)}, " ")
}
