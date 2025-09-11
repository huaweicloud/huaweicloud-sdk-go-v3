package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbOmInstanceResponse Response Object
type DeleteDbOmInstanceResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbOmInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbOmInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbOmInstanceResponse", string(data)}, " ")
}
