package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAddOrDeleteTagOnClusterReq struct {

	// 操作类型。 说明：通过该属性标识当前所需的操作类型。 - create：批量添加标签。 - delete：批量删除标签。
	Action string `json:"action"`

	// 标签列表。
	Tags []Tag `json:"tags"`

	// 标签列表。
	SysTags *[]SysTags `json:"sysTags,omitempty"`
}

func (o BatchAddOrDeleteTagOnClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagOnClusterReq struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagOnClusterReq", string(data)}, " ")
}
