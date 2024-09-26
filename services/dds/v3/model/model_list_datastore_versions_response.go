package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatastoreVersionsResponse Response Object
type ListDatastoreVersionsResponse struct {

	// 数据库版本。支持3.4及以上版本。
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatastoreVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoreVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListDatastoreVersionsResponse", string(data)}, " ")
}
