package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSecrecyLevelDto struct {

	// 密级id列表，密级id可以通过查询接口获取。
	Ids *[]string `json:"ids,omitempty"`
}

func (o BatchDeleteSecrecyLevelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecrecyLevelDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecrecyLevelDto", string(data)}, " ")
}
