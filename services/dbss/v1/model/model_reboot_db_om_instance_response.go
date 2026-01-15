package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootDbOmInstanceResponse Response Object
type RebootDbOmInstanceResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebootDbOmInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootDbOmInstanceResponse struct{}"
	}

	return strings.Join([]string{"RebootDbOmInstanceResponse", string(data)}, " ")
}
