package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolsRequestBody This is a auto create Body Object
type BatchDeletePoolsRequestBody struct {

	// 待删除的后端服务器组id列表。
	Pools []string `json:"pools"`
}

func (o BatchDeletePoolsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolsRequestBody", string(data)}, " ")
}
