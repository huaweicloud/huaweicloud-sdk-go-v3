package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReduceCnRequestBody struct {

	// **参数解释**： 缩容节点ID集合。 **约束限制**： 不涉及。
	NodeIdList []string `json:"node_id_list"`
}

func (o ReduceCnRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReduceCnRequestBody struct{}"
	}

	return strings.Join([]string{"ReduceCnRequestBody", string(data)}, " ")
}
