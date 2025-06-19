package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBlacklistDeleteDto 包含删除IP黑名单的参数，目前只会指定生效范围进行删除。
type IpBlacklistDeleteDto struct {

	// 指定生效范围，1指定生效范围为EIP进行删除,2指定生效范围为NAT进行删除，1,2生效范围为EIP和NAT进行删除
	EffectScope *[]int32 `json:"effect_scope,omitempty"`
}

func (o IpBlacklistDeleteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBlacklistDeleteDto struct{}"
	}

	return strings.Join([]string{"IpBlacklistDeleteDto", string(data)}, " ")
}
