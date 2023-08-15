package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteResourceGroupsRequestBody struct {

	// 需要批量删除的资源分组ID列表
	GroupIds []string `json:"group_ids"`
}

func (o BatchDeleteResourceGroupsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceGroupsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceGroupsRequestBody", string(data)}, " ")
}
