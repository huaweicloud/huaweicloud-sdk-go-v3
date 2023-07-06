package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHighRiskCommandsResponse Response Object
type UpdateHighRiskCommandsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHighRiskCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHighRiskCommandsResponse struct{}"
	}

	return strings.Join([]string{"UpdateHighRiskCommandsResponse", string(data)}, " ")
}
