package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartClusterResponse Response Object
type RestartClusterResponse struct {

	// 作业ID
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartClusterResponse struct{}"
	}

	return strings.Join([]string{"RestartClusterResponse", string(data)}, " ")
}
