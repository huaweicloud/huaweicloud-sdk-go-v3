package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobByIdRequest struct {

	// 作业ID
	JobId string `json:"job_id" xml:"job_id"`
}

func (o ShowJobByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowJobByIdRequest", string(data)}, " ")
}
