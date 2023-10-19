package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagCentralNetworkRequestBody 创建标签的请求体。
type UntagCentralNetworkRequestBody struct {

	// 包含标签。
	Tags []Tag `json:"tags"`
}

func (o UntagCentralNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagCentralNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"UntagCentralNetworkRequestBody", string(data)}, " ")
}
