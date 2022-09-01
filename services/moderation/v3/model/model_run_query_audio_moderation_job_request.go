package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunQueryAudioModerationJobRequest struct {
	JobId string `json:"job_id" xml:"job_id"`
}

func (o RunQueryAudioModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryAudioModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunQueryAudioModerationJobRequest", string(data)}, " ")
}
