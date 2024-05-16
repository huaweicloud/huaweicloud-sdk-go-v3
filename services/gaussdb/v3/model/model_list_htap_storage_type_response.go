package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapStorageTypeResponse Response Object
type ListHtapStorageTypeResponse struct {

	// 存储类型。
	StorageType    *[]HtapStorageTypeStorageType `json:"storage_type,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListHtapStorageTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapStorageTypeResponse struct{}"
	}

	return strings.Join([]string{"ListHtapStorageTypeResponse", string(data)}, " ")
}
