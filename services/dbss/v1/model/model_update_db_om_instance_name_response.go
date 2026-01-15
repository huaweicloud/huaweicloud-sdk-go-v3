package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDbOmInstanceNameResponse Response Object
type UpdateDbOmInstanceNameResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDbOmInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbOmInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDbOmInstanceNameResponse", string(data)}, " ")
}
