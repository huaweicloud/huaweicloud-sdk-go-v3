package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicationResponse Response Object
type DeletePublicationResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicationResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicationResponse", string(data)}, " ")
}
