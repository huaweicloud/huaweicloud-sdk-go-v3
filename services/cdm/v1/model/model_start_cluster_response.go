package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartClusterResponse struct {
	// 作业ID

	JobId          *[]string `json:"jobId,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o StartClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartClusterResponse struct{}"
	}

	return strings.Join([]string{"StartClusterResponse", string(data)}, " ")
}
