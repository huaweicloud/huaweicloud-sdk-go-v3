package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReduceInstanceNodeRequestBody struct {

	// 删除的节点数量。
	Num *int32 `json:"num,omitempty"`

	// 指定删除节点的ID列表。 - num与node_list必须有一个字段传值 - 如果num与node_list同时传值时，则以node_list的值为主 - 删除的节点角色不能是Primary - 如果是多AZ实例，请确保删除节点后，每个AZ至少保留一个节点
	NodeList *[]string `json:"node_list,omitempty"`
}

func (o ReduceInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ReduceInstanceNodeRequestBody", string(data)}, " ")
}
