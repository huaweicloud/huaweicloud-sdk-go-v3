package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeCbhRequestBody 云堡垒机实例请求对象。
type UpgradeCbhRequestBody struct {

	// 实例id
	ServerId string `json:"server_id"`

	// 定时升级的时间，需要比当前时间大24小时
	UpgradeTime int64 `json:"upgrade_time"`

	// 是否取消升级定时任务，已开始任务不可取消。 - true：取消 - false：无影响
	Cancel *bool `json:"cancel,omitempty"`
}

func (o UpgradeCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeCbhRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeCbhRequestBody", string(data)}, " ")
}
