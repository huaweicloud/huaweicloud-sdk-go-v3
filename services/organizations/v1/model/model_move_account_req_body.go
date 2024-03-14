package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAccountReqBody MoveAccount 操作的请求体。
type MoveAccountReqBody struct {

	// 要移出账号的根或组织单元的唯一标识符（ID）。
	SourceParentId string `json:"source_parent_id"`

	// 要移入账号的根或组织单元的唯一标识符（ID）。
	DestinationParentId string `json:"destination_parent_id"`
}

func (o MoveAccountReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAccountReqBody struct{}"
	}

	return strings.Join([]string{"MoveAccountReqBody", string(data)}, " ")
}
