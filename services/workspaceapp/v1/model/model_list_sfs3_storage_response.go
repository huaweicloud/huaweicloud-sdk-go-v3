package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSfs3StorageResponse Response Object
type ListSfs3StorageResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 存储列表。
	Items          *[]Sfs3Storage `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSfs3StorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSfs3StorageResponse struct{}"
	}

	return strings.Join([]string{"ListSfs3StorageResponse", string(data)}, " ")
}
