package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableType 表类型。
type TableType struct {
}

func (o TableType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableType struct{}"
	}

	return strings.Join([]string{"TableType", string(data)}, " ")
}
