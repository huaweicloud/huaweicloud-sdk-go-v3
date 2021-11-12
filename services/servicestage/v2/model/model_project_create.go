package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectCreate struct {
	// 项目名称。

	Name string `json:"name"`
}

func (o ProjectCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectCreate struct{}"
	}

	return strings.Join([]string{"ProjectCreate", string(data)}, " ")
}
