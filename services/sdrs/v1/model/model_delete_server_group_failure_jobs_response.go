package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteServerGroupFailureJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupFailureJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupFailureJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupFailureJobsResponse", string(data)}, " ")
}
