package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoEnlargePolicyResponse Response Object
type ShowAutoEnlargePolicyResponse struct {

	// 自动扩容开关。 - on：开启磁盘自动扩容策略。 - off: 关闭磁盘自动扩容策略。 默认值为on。
	SwitchOption *string `json:"switch_option,omitempty"`

	Policy         *DiskAutoExpansionPolicy `json:"policy,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoEnlargePolicyResponse", string(data)}, " ")
}
