package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchInferDeploymentVersionResponse Response Object
type SwitchInferDeploymentVersionResponse struct {

	// **参数解释：** 服务部署id **取值范围：** 不涉及。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchInferDeploymentVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchInferDeploymentVersionResponse struct{}"
	}

	return strings.Join([]string{"SwitchInferDeploymentVersionResponse", string(data)}, " ")
}
