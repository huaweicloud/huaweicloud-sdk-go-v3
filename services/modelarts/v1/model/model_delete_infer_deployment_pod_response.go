package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferDeploymentPodResponse Response Object
type DeleteInferDeploymentPodResponse struct {

	// **参数解释：** pod名字。 **取值范围：** 不涉及。
	PodName        *string `json:"pod_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInferDeploymentPodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferDeploymentPodResponse struct{}"
	}

	return strings.Join([]string{"DeleteInferDeploymentPodResponse", string(data)}, " ")
}
