package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDatabaseInfoResponse struct {
	// DDM实例id。

	InstanceId *string `json:"instanceId,omitempty"`
	// 任务ID。

	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDatabaseInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseInfoResponse", string(data)}, " ")
}
