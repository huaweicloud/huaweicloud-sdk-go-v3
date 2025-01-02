package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DelayRestoreTable struct {

	// 表名称
	Name string `json:"name"`
}

func (o DelayRestoreTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelayRestoreTable struct{}"
	}

	return strings.Join([]string{"DelayRestoreTable", string(data)}, " ")
}
