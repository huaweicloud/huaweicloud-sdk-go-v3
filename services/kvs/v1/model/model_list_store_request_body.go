package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListStoreRequestBody struct {

	// 上次返回的游标位置，本次响应包含该仓名。 - 长度：[16,52] - 取值字符限制：[a-z0-9-]+
	CursorName *string `bson:"cursor_name,omitempty"`

	// 响应返回的仓个数。
	Limit *int32 `bson:"limit,omitempty"`
}

func (o ListStoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoreRequestBody struct{}"
	}

	return strings.Join([]string{"ListStoreRequestBody", string(data)}, " ")
}
