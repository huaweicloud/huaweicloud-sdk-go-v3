package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopClusterResponse Response Object
type StopClusterResponse struct {

	// 作业ID
	JobId          *[]string `json:"jobId,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o StopClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopClusterResponse struct{}"
	}

	return strings.Join([]string{"StopClusterResponse", string(data)}, " ")
}
