package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageGearsResponse Response Object
type ListStorageGearsResponse struct {

	// 支持的存储档位列表
	StorageGears *[]StorageGearV2 `json:"storage_gears,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStorageGearsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageGearsResponse struct{}"
	}

	return strings.Join([]string{"ListStorageGearsResponse", string(data)}, " ")
}
