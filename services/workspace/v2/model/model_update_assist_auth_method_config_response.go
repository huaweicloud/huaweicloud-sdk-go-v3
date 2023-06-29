package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssistAuthMethodConfigResponse Response Object
type UpdateAssistAuthMethodConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAssistAuthMethodConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssistAuthMethodConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssistAuthMethodConfigResponse", string(data)}, " ")
}
