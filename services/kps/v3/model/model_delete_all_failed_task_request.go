package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAllFailedTaskRequest struct {
}

func (o DeleteAllFailedTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllFailedTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteAllFailedTaskRequest", string(data)}, " ")
}
