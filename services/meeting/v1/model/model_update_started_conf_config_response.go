package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStartedConfConfigResponse Response Object
type UpdateStartedConfConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStartedConfConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartedConfConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateStartedConfConfigResponse", string(data)}, " ")
}
