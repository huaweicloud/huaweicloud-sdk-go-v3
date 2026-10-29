package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTaskResponse Response Object
type RunTaskResponse struct {

	// **参数解释:** 1次代码检查对应的执行记录ID。 **取值范围:** 32位UUID。
	ExecId         *string `json:"exec_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTaskResponse struct{}"
	}

	return strings.Join([]string{"RunTaskResponse", string(data)}, " ")
}
