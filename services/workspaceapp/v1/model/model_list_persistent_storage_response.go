package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersistentStorageResponse Response Object
type ListPersistentStorageResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 存储列表。
	Items          *[]PersistentStorage `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListPersistentStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersistentStorageResponse struct{}"
	}

	return strings.Join([]string{"ListPersistentStorageResponse", string(data)}, " ")
}
