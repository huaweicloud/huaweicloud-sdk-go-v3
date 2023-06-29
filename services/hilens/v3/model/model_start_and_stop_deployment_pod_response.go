package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAndStopDeploymentPodResponse Response Object
type StartAndStopDeploymentPodResponse struct {

	// pod的ID
	PodId          *string `json:"pod_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAndStopDeploymentPodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAndStopDeploymentPodResponse struct{}"
	}

	return strings.Join([]string{"StartAndStopDeploymentPodResponse", string(data)}, " ")
}
