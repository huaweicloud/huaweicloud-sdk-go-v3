package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountReqBody UpdateAccountReqBody 操作的请求体。
type UpdateAccountReqBody struct {

	// 描述信息。如果为空字符串，则表示把描述信息设置为空。如果为NULL，则不做任何处理。
	Description *string `json:"description,omitempty"`
}

func (o UpdateAccountReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountReqBody struct{}"
	}

	return strings.Join([]string{"UpdateAccountReqBody", string(data)}, " ")
}
