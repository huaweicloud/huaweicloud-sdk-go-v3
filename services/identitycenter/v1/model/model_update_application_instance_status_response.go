package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceStatusResponse Response Object
type UpdateApplicationInstanceStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationInstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceStatusResponse", string(data)}, " ")
}
