package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeDatabasesSingleInstance struct {

	// 实例当前的内核版本。
	CurrentVersion string `json:"current_version"`

	// 实例id。
	InstanceId string `json:"instance_id"`
}

func (o UpgradeDatabasesSingleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabasesSingleInstance struct{}"
	}

	return strings.Join([]string{"UpgradeDatabasesSingleInstance", string(data)}, " ")
}
