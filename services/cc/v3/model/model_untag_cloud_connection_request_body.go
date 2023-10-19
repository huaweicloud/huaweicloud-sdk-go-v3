package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagCloudConnectionRequestBody 删除云连接实例标签的请求体。
type UntagCloudConnectionRequestBody struct {

	// 包含标签。
	Tags []Tag `json:"tags"`
}

func (o UntagCloudConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagCloudConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UntagCloudConnectionRequestBody", string(data)}, " ")
}
