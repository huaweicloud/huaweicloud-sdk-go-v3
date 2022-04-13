package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEnvironmentRequest struct {
	// 环境ID。

	EnvironmentId string `json:"environment_id"`
}

func (o DeleteEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentRequest", string(data)}, " ")
}
