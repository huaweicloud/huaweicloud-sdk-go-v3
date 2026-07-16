package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpGet struct {

	// **参数解释**： http获取指标的url路径，与下面的端口必须同时填或者不填。 **取值范围**： 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释**： http获取指标的端口，与上面的url路径必须同时填或者不填。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`
}

func (o HttpGet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpGet struct{}"
	}

	return strings.Join([]string{"HttpGet", string(data)}, " ")
}
