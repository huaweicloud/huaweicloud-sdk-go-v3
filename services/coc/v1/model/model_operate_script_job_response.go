package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateScriptJobResponse Response Object
type OperateScriptJobResponse struct {

	// 操作的工单execute_uuid
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o OperateScriptJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateScriptJobResponse struct{}"
	}

	return strings.Join([]string{"OperateScriptJobResponse", string(data)}, " ")
}
