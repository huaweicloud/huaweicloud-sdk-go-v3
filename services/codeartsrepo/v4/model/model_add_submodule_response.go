package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSubmoduleResponse Response Object
type AddSubmoduleResponse struct {

	// **参数解释：** 创建成功提交ID。
	Result *string `json:"result,omitempty"`

	// **参数解释：** 创建返回状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddSubmoduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubmoduleResponse struct{}"
	}

	return strings.Join([]string{"AddSubmoduleResponse", string(data)}, " ")
}
