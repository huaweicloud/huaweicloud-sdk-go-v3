package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectWorkHoursRequest Request Object
type ShowProjectWorkHoursRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *ShowProjectWorkHoursRequestBody `json:"body,omitempty"`
}

func (o ShowProjectWorkHoursRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWorkHoursRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectWorkHoursRequest", string(data)}, " ")
}
