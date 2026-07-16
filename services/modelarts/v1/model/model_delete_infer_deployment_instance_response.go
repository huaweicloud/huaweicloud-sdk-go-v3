package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferDeploymentInstanceResponse Response Object
type DeleteInferDeploymentInstanceResponse struct {

	// **参数解释：** 服务实例名字。 **取值范围：** 不涉及。
	InstanceName   *string `json:"instance_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInferDeploymentInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferDeploymentInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInferDeploymentInstanceResponse", string(data)}, " ")
}
