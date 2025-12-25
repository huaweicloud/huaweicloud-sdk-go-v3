package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableAlias 表别名
type TableAlias struct {
}

func (o TableAlias) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableAlias struct{}"
	}

	return strings.Join([]string{"TableAlias", string(data)}, " ")
}
