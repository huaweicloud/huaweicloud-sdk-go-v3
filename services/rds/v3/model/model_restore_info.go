package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInfo struct {

	// 需恢复的原库名称
	OldName string `json:"old_name"`

	// 恢复后的新库名称
	NewName string `json:"new_name"`
}

func (o RestoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInfo struct{}"
	}

	return strings.Join([]string{"RestoreInfo", string(data)}, " ")
}
