package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRedistributionResponse Response Object
type RestoreRedistributionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreRedistributionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedistributionResponse struct{}"
	}

	return strings.Join([]string{"RestoreRedistributionResponse", string(data)}, " ")
}
