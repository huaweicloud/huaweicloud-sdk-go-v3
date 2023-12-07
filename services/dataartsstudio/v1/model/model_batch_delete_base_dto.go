package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteBaseDto struct {

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o BatchDeleteBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBaseDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaseDto", string(data)}, " ")
}
