package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SoarBaseRsp 请求的返回对象
type SoarBaseRsp struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**: 是否成功 **取值范围**: - true  成功 - false 失败
	Success *bool `json:"success,omitempty"`

	// **参数解释**: 请求的ID **约束限制**: 不涉及
	RequestId *string `json:"request_id,omitempty"`
}

func (o SoarBaseRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SoarBaseRsp struct{}"
	}

	return strings.Join([]string{"SoarBaseRsp", string(data)}, " ")
}
