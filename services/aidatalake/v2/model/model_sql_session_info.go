package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlSessionInfo struct {

	// **参数解释**：Session的ID。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	SessionId *string `json:"session_id,omitempty"`

	Status *SqlSessionStatus `json:"status,omitempty"`

	// **参数解释**：session创建时间。 **取值范围**：不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：端点信息。
	EndpointInfo *interface{} `json:"endpoint_info,omitempty"`
}

func (o SqlSessionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlSessionInfo struct{}"
	}

	return strings.Join([]string{"SqlSessionInfo", string(data)}, " ")
}
