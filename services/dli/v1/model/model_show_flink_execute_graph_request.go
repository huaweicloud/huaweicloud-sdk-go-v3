package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkExecuteGraphRequest Request Object
type ShowFlinkExecuteGraphRequest struct {

	// 作业ID。
	JobId int64 `json:"job_id"`
}

func (o ShowFlinkExecuteGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkExecuteGraphRequest struct{}"
	}

	return strings.Join([]string{"ShowFlinkExecuteGraphRequest", string(data)}, " ")
}
