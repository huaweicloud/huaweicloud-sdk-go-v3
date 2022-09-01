package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteApplicationConfigurationRequest struct {

	// 应用ID。
	ApplicationId string `json:"application_id" xml:"application_id"`

	// 环境ID。
	EnvironmentId string `json:"environment_id" xml:"environment_id"`
}

func (o DeleteApplicationConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationConfigurationRequest", string(data)}, " ")
}
