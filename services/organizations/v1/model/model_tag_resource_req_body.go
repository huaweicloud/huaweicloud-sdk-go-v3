package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceReqBody TagResource 操作的请求体。
type TagResourceReqBody struct {

	// 要添加到指定资源的标签列表。
	Tags []TagDto `json:"tags"`
}

func (o TagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceReqBody struct{}"
	}

	return strings.Join([]string{"TagResourceReqBody", string(data)}, " ")
}
