package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteConnectionNewRequestBody struct {

	// 是否删除所有连接
	DeleteAll *bool `json:"delete_all,omitempty"`

	// 连接ID列表
	ConnectionIds *[]ConnectionIdsItem `json:"connection_ids,omitempty"`
}

func (o BatchDeleteConnectionNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConnectionNewRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteConnectionNewRequestBody", string(data)}, " ")
}
