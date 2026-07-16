package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferDeploymentVersionResponse Response Object
type DeleteInferDeploymentVersionResponse struct {

	// **参数解释：** 服务部署id **取值范围：** 不涉及。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInferDeploymentVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferDeploymentVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteInferDeploymentVersionResponse", string(data)}, " ")
}
