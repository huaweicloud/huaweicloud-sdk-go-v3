package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDatabaseInfoResponse struct {

	// DDM实例id。
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId"`

	// 任务ID。
	JobId          *string `json:"jobId,omitempty" xml:"jobId"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDatabaseInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseInfoResponse", string(data)}, " ")
}
