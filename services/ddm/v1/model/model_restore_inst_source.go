package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstSource struct {

	// 恢复时间。
	RestoreTime float32 `json:"restore_time"`
}

func (o RestoreInstSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstSource struct{}"
	}

	return strings.Join([]string{"RestoreInstSource", string(data)}, " ")
}
