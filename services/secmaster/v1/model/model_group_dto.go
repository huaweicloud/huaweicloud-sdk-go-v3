package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupDto struct {

	// 分组名称
	Name string `json:"name"`
}

func (o GroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupDto struct{}"
	}

	return strings.Join([]string{"GroupDto", string(data)}, " ")
}
