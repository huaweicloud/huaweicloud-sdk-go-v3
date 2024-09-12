package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoreResponse Response Object
type ListStoreResponse struct {

	// 本次响应后的游标位置，下次请求时携带。 - 长度：[16,52] - 取值字符限制：[a-z0-9-]+ > 如果为空，表示后面无更多仓名。
	CursorName *string `bson:"cursor_name,omitempty"`

	// 返回的仓名列表。
	Stores         *[]string `bson:"stores,omitempty"`
	HttpStatusCode int       `bson:"-"`
}

func (o ListStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoreResponse struct{}"
	}

	return strings.Join([]string{"ListStoreResponse", string(data)}, " ")
}
