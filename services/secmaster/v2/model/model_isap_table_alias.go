package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTableAlias 表别名
type IsapTableAlias struct {
}

func (o IsapTableAlias) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableAlias struct{}"
	}

	return strings.Join([]string{"IsapTableAlias", string(data)}, " ")
}
