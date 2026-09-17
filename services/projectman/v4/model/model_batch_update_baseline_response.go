package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateBaselineResponse Response Object
type BatchUpdateBaselineResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： - success：批量操作成功 - error：批量操作失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 返回消息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	Result         *BatchResultVo `json:"result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchUpdateBaselineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBaselineResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateBaselineResponse", string(data)}, " ")
}
