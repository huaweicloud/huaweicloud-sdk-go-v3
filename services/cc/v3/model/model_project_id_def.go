package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectIdDef 实例所属项目ID。
type ProjectIdDef struct {
}

func (o ProjectIdDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectIdDef struct{}"
	}

	return strings.Join([]string{"ProjectIdDef", string(data)}, " ")
}
