package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectInfoCoverAreaInfo 防护覆盖面信息
type ProtectInfoCoverAreaInfo struct {

	// **参数解释**: 防护主机数量 **取值范围**: 最小值0，最大值2147483647
	ProtectHostNum *int32 `json:"protect_host_num,omitempty"`

	// **参数解释**: 未防护主机数量 **取值范围**: 最小值0，最大值2147483647
	UnProtectHostNum *int32 `json:"un_protect_host_num,omitempty"`

	// **参数解释**: 防护率 **取值范围**: 最小值0，最大值1
	ProtectRate *float32 `json:"protect_rate,omitempty"`

	// **参数解释**: 击败用户率 **取值范围**: 最小值0，最大值1
	BeatRate *float32 `json:"beat_rate,omitempty"`
}

func (o ProtectInfoCoverAreaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectInfoCoverAreaInfo struct{}"
	}

	return strings.Join([]string{"ProtectInfoCoverAreaInfo", string(data)}, " ")
}
