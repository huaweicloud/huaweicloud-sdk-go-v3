package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartDdmInstanceResponse Response Object
type RestartDdmInstanceResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartDdmInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartDdmInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartDdmInstanceResponse", string(data)}, " ")
}
