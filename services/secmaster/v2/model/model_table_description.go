package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableDescription 表描述
type TableDescription struct {
}

func (o TableDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableDescription struct{}"
	}

	return strings.Join([]string{"TableDescription", string(data)}, " ")
}
