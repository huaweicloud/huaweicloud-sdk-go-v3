package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindDbOmEipResponse Response Object
type UnbindDbOmEipResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindDbOmEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindDbOmEipResponse struct{}"
	}

	return strings.Join([]string{"UnbindDbOmEipResponse", string(data)}, " ")
}
