package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RollbackInfo struct {

	// 回滚的目标云硬盘UUID。
	VolumeId string `json:"volume_id" xml:"volume_id"`
}

func (o RollbackInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackInfo struct{}"
	}

	return strings.Join([]string{"RollbackInfo", string(data)}, " ")
}
