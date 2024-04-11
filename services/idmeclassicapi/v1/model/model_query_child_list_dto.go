package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryChildListDto struct {

	// 父节点实例ID。
	ParentId *string `json:"parentId,omitempty"`
}

func (o QueryChildListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryChildListDto struct{}"
	}

	return strings.Join([]string{"QueryChildListDto", string(data)}, " ")
}
