package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiVirusSwitchDto struct {

	// 反病毒开关状态
	AntiVirusStatus *int32 `json:"anti_virus_status,omitempty"`

	// 防护对象ID
	ObjectId *string `json:"object_id,omitempty"`
}

func (o AntiVirusSwitchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusSwitchDto struct{}"
	}

	return strings.Join([]string{"AntiVirusSwitchDto", string(data)}, " ")
}
