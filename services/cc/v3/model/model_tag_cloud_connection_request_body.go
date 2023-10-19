package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagCloudConnectionRequestBody 创建云连接实例标签的请求体。
type TagCloudConnectionRequestBody struct {

	// 包含标签。
	Tags []Tag `json:"tags"`
}

func (o TagCloudConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCloudConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"TagCloudConnectionRequestBody", string(data)}, " ")
}
