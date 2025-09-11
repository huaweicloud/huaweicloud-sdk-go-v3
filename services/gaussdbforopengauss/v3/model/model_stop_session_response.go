package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSessionResponse Response Object
type StopSessionResponse struct {

	// **参数解释**： 查杀指定会话ID列表。
	SessionIds *[]string `json:"session_ids,omitempty"`

	// **参数解释**： 结束会话操作执行是否成功。 **取值范围**: - true：成功。 - false：失败。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StopSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSessionResponse struct{}"
	}

	return strings.Join([]string{"StopSessionResponse", string(data)}, " ")
}
