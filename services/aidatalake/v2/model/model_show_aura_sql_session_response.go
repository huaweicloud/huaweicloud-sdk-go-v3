package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraSqlSessionResponse Response Object
type ShowAuraSqlSessionResponse struct {

	// **参数解释**：Session的ID。 **取值范围**：长度为1~32个字符，支持大小写英文字母、数字、连字符。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：session的状态。 **取值范围**：   - RUNNING：运行中。   - CLOSED：已关闭。   - WAITING：等待中。   - CREATING：创建中。   - FAIL：失败。
	Status *string `json:"status,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAuraSqlSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraSqlSessionResponse struct{}"
	}

	return strings.Join([]string{"ShowAuraSqlSessionResponse", string(data)}, " ")
}
