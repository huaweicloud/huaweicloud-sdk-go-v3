package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSceneByPageResponse Response Object
type ShowSceneByPageResponse struct {
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoScene `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSceneByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSceneByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowSceneByPageResponse", string(data)}, " ")
}
