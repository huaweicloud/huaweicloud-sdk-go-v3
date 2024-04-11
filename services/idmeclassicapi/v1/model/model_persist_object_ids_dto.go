package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsDto struct {

	// 数据实例ID列表。
	Ids []string `json:"ids"`
}

func (o PersistObjectIdsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsDto", string(data)}, " ")
}
