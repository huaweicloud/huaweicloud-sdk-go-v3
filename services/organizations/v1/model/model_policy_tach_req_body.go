package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyTachReqBody AttachPolicy 和 DetachPolicy 操作的请求体。
type PolicyTachReqBody struct {

	// 根、组织单元或账号的唯一标识符（ID）。
	EntityId string `json:"entity_id"`
}

func (o PolicyTachReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTachReqBody struct{}"
	}

	return strings.Join([]string{"PolicyTachReqBody", string(data)}, " ")
}
