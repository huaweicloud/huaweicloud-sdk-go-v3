package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestpointByPageResponse Response Object
type ShowTestpointByPageResponse struct {
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoTestPoint `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTestpointByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestpointByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowTestpointByPageResponse", string(data)}, " ")
}
