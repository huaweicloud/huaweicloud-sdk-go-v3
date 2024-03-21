package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdDto struct {

	// 唯一标识。
	Id string `json:"id"`
}

func (o PersistObjectIdDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdDto", string(data)}, " ")
}
