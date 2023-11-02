package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPositionResultResponse Response Object
type ShowPositionResultResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 位点信息
	Position       *string `json:"position,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPositionResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPositionResultResponse struct{}"
	}

	return strings.Join([]string{"ShowPositionResultResponse", string(data)}, " ")
}
