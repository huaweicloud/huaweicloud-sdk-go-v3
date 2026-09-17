package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeSlowLogArchiveRequest Request Object
type InvokeSlowLogArchiveRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o InvokeSlowLogArchiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeSlowLogArchiveRequest struct{}"
	}

	return strings.Join([]string{"InvokeSlowLogArchiveRequest", string(data)}, " ")
}
