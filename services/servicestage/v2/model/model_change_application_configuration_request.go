package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeApplicationConfigurationRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`

	Body *ApplicationConfigModify `json:"body,omitempty"`
}

func (o ChangeApplicationConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeApplicationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ChangeApplicationConfigurationRequest", string(data)}, " ")
}
