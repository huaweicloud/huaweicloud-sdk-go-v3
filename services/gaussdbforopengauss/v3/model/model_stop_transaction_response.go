package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopTransactionResponse Response Object
type StopTransactionResponse struct {

	// **参数解释**: 查杀事务ID列表。
	SessionIds *[]string `json:"session_ids,omitempty"`

	// **参数解释**: 结束事务操作执行是否成功。 **取值范围**: -true：成功。 -false：未成功。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StopTransactionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTransactionResponse struct{}"
	}

	return strings.Join([]string{"StopTransactionResponse", string(data)}, " ")
}
