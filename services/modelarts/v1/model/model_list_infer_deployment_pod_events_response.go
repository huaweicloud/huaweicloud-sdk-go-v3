package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferDeploymentPodEventsResponse Response Object
type ListInferDeploymentPodEventsResponse struct {

	// **参数解释：** 服务Pod事件列表。
	Body           *[]ServicePodEventResponse `json:"body,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListInferDeploymentPodEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferDeploymentPodEventsResponse struct{}"
	}

	return strings.Join([]string{"ListInferDeploymentPodEventsResponse", string(data)}, " ")
}
