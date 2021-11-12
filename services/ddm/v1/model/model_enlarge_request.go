package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto request Object
type EnlargeRequest struct {
	// 当前进行节点扩容的DDM实例底层虚机规格id

	FlavorId string `json:"flavor_id"`
	// 需要扩容的节点个数

	NodeNumber int32 `json:"node_number"`
}

func (o EnlargeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeRequest struct{}"
	}

	return strings.Join([]string{"EnlargeRequest", string(data)}, " ")
}
