package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopDbOmInstanceResponse Response Object
type StopDbOmInstanceResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopDbOmInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDbOmInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopDbOmInstanceResponse", string(data)}, " ")
}
