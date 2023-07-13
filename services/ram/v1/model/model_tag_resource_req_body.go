package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceReqBody The request body of the TagResource operation.
type TagResourceReqBody struct {

	// 一个或多个标签键值对的列表。标签键必须存在，而不是空字符串。标签值必须存在，但可以是空字符串。
	Tags []Tag `json:"tags"`
}

func (o TagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceReqBody struct{}"
	}

	return strings.Join([]string{"TagResourceReqBody", string(data)}, " ")
}
