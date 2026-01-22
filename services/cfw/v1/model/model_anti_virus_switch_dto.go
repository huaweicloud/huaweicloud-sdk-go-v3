package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiVirusSwitchDto struct {

	// 参数解释： 反病毒开关状态，为必传参数 约束限制： 不涉及 取值范围： 0：开启 1：关闭 默认取值： 不涉及
	AntiVirusStatus *int32 `json:"anti_virus_status,omitempty"`

	// 防护对象ID，为必传参数
	ObjectId *string `json:"object_id,omitempty"`
}

func (o AntiVirusSwitchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusSwitchDto struct{}"
	}

	return strings.Join([]string{"AntiVirusSwitchDto", string(data)}, " ")
}
