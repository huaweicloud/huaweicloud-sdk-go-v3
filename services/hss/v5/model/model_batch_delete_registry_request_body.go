package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRegistryRequestBody 镜像仓批量删除信息
type BatchDeleteRegistryRequestBody struct {

	// 镜像仓ID列表
	Ids []string `json:"ids"`
}

func (o BatchDeleteRegistryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRegistryRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteRegistryRequestBody", string(data)}, " ")
}
