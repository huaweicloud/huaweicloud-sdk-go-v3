package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrozenInfo 资源的冻结信息，当资源被冻结时返回冻结的类型信息。
type FrozenInfo struct {

	// 冻结场景。可选值如下： - ARREAR：欠费冻结 - POLICE：公安冻结 - ILLEGAL：违规冻结
	Scene *string `json:"scene,omitempty"`

	// 冻结后的影响。可选值如下： - 1：冻结后可释放 - 2：冻结后不可释放
	Effect *int32 `json:"effect,omitempty"`
}

func (o FrozenInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrozenInfo struct{}"
	}

	return strings.Join([]string{"FrozenInfo", string(data)}, " ")
}
