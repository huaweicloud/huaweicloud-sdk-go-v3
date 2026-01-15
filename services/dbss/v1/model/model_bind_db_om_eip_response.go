package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindDbOmEipResponse Response Object
type BindDbOmEipResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindDbOmEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDbOmEipResponse struct{}"
	}

	return strings.Join([]string{"BindDbOmEipResponse", string(data)}, " ")
}
