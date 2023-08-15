package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeDatabaseVersionRequestBody struct {

	// 升级模式。  取值为“minimized_interrupt_time”为中断时间最短优先模式：升级过程对业务影响相对较小的升级方式  取值为“minimized_upgrade_time”为升级时长最短优先模式：升级过程时长相对较快的升级方式。  默认取值为“minimized_interrupt_time”。
	UpgradeMode *string `json:"upgrade_mode,omitempty"`
}

func (o UpgradeDatabaseVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabaseVersionRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeDatabaseVersionRequestBody", string(data)}, " ")
}
