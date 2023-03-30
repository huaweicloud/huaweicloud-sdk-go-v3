package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// The request body of the UntagResource operation.
type UntagResourceReqBody struct {

	// 一个或多个标签键值对的列表。标签键必须存在，而不是空字符串。标签值必须存在，但可以是空字符串。
	Tags []Untag `json:"tags"`
}

func (o UntagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagResourceReqBody struct{}"
	}

	return strings.Join([]string{"UntagResourceReqBody", string(data)}, " ")
}
