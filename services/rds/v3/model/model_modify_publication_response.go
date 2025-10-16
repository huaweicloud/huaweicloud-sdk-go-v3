package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPublicationResponse Response Object
type ModifyPublicationResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyPublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPublicationResponse struct{}"
	}

	return strings.Join([]string{"ModifyPublicationResponse", string(data)}, " ")
}
