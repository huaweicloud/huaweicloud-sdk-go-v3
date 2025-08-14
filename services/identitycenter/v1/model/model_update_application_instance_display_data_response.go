package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceDisplayDataResponse Response Object
type UpdateApplicationInstanceDisplayDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationInstanceDisplayDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceDisplayDataResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceDisplayDataResponse", string(data)}, " ")
}
