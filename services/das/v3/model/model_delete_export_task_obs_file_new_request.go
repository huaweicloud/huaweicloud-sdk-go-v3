package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskObsFileNewRequest Request Object
type DeleteExportTaskObsFileNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	Id int64 `json:"id"`
}

func (o DeleteExportTaskObsFileNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskObsFileNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskObsFileNewRequest", string(data)}, " ")
}
