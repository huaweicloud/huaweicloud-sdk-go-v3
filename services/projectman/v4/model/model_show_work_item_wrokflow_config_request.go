package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkItemWrokflowConfigRequest Request Object
type ShowWorkItemWrokflowConfigRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 看板id
	BoardId string `json:"board_id"`
}

func (o ShowWorkItemWrokflowConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkItemWrokflowConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkItemWrokflowConfigRequest", string(data)}, " ")
}
