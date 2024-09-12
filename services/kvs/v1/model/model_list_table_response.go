package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableResponse Response Object
type ListTableResponse struct {

	// 本次响应后的游标位置，下次请求时携带，如果为空，表示后面无更多。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+ > 如果为空，表示后面无更多。
	CursorName *string `bson:"cursor_name,omitempty"`

	// 返回的表名列表。 - 长度：最大100
	TableNames     *[]string `bson:"table_names,omitempty"`
	HttpStatusCode int       `bson:"-"`
}

func (o ListTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableResponse struct{}"
	}

	return strings.Join([]string{"ListTableResponse", string(data)}, " ")
}
