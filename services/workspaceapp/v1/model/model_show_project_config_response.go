package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectConfigResponse Response Object
type ShowProjectConfigResponse struct {
	ProjectConfig  *ProjectConfig `json:"project_config,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowProjectConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectConfigResponse", string(data)}, " ")
}
