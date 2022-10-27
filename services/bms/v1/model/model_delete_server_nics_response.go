package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteServerNicsResponse struct {

	//
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServerNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerNicsResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerNicsResponse", string(data)}, " ")
}
