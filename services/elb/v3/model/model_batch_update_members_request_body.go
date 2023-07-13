package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateMembersRequestBody This is a auto create Body Object
type BatchUpdateMembersRequestBody struct {

	// 后端服务器对象。
	Members []BatchUpdateMembersOption `json:"members"`
}

func (o BatchUpdateMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersRequestBody", string(data)}, " ")
}
