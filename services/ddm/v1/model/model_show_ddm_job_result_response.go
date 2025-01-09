package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmJobResultResponse Response Object
type ShowDdmJobResultResponse struct {
	Job            *JobInfo `json:"job,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowDdmJobResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmJobResultResponse struct{}"
	}

	return strings.Join([]string{"ShowDdmJobResultResponse", string(data)}, " ")
}
