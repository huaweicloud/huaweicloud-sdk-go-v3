package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto request Object
type EnlargeRequest struct {
	// 当前进行节点扩容的DDM实例底层虚机规格id

	FlavorId string `json:"flavor_id"`
	// 需要扩容的节点个数

	NodeNumber int32 `json:"node_number"`
	// 组id，指定当前进行节点扩容的组。当实例的组>1时，必填。

	GroupId *string `json:"group_id,omitempty"`
}

func (o EnlargeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeRequest struct{}"
	}

	return strings.Join([]string{"EnlargeRequest", string(data)}, " ")
}
