package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCocIncidentDetailRequest Request Object
type ShowCocIncidentDetailRequest struct {

	// 事件单号
	IncidentNum string `json:"incident_num"`
}

func (o ShowCocIncidentDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCocIncidentDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCocIncidentDetailRequest", string(data)}, " ")
}
