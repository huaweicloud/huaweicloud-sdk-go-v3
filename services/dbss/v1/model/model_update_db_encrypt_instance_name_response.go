package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDbEncryptInstanceNameResponse Response Object
type UpdateDbEncryptInstanceNameResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDbEncryptInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbEncryptInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDbEncryptInstanceNameResponse", string(data)}, " ")
}
