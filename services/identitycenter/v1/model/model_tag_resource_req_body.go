package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceReqBody the request body of TagResource
type TagResourceReqBody struct {

	// 用于管理资源的一组键值对
	Tags []TagDto `json:"tags"`
}

func (o TagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceReqBody struct{}"
	}

	return strings.Join([]string{"TagResourceReqBody", string(data)}, " ")
}
