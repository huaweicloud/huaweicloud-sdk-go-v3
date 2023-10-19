package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionDef 版本标识符
type VersionDef struct {
}

func (o VersionDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDef struct{}"
	}

	return strings.Join([]string{"VersionDef", string(data)}, " ")
}
