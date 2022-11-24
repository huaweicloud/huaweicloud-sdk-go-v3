package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotebookToolDto struct {

	// 显示名称
	DisplayName string `json:"display_name"`

	Profile *Profile `json:"profile"`
}

func (o NotebookToolDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookToolDto struct{}"
	}

	return strings.Join([]string{"NotebookToolDto", string(data)}, " ")
}
