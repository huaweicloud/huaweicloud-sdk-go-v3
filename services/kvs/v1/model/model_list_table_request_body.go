package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTableRequestBody struct {

	// 上次返回游标位置，本次响应包含该table，空表示遍历完。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+ > 如果为空，表示后面无更多。
	CursorName *string `bson:"cursor_name,omitempty"`

	// 响应返回的表个数。 - 长度：最大100
	Limit *int32 `bson:"limit,omitempty"`
}

func (o ListTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableRequestBody struct{}"
	}

	return strings.Join([]string{"ListTableRequestBody", string(data)}, " ")
}
