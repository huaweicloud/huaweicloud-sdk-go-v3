package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareFolderResponse Response Object
type ListShareFolderResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 存储声明。
	Items          *[]SharePersistentStorageClaim `json:"items,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListShareFolderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareFolderResponse struct{}"
	}

	return strings.Join([]string{"ListShareFolderResponse", string(data)}, " ")
}
