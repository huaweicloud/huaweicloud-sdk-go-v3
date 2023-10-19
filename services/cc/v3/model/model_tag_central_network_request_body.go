package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagCentralNetworkRequestBody 创建标签的请求体。
type TagCentralNetworkRequestBody struct {

	// 包含标签。
	Tags []Tag `json:"tags"`
}

func (o TagCentralNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCentralNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"TagCentralNetworkRequestBody", string(data)}, " ")
}
