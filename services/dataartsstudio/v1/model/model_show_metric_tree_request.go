package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricTreeRequest Request Object
type ShowMetricTreeRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowMetricTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricTreeRequest", string(data)}, " ")
}
