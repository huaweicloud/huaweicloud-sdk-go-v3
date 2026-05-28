package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreInstTarget struct {

	// 实例id。
	InstanceId string `json:"instance_id"`
}

func (o RestoreInstTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstTarget struct{}"
	}

	return strings.Join([]string{"RestoreInstTarget", string(data)}, " ")
}
