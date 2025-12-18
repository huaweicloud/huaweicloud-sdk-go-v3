package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableResourceResponse Response Object
type ShowAvailableResourceResponse struct {
	AvailableResource *[]AvailableResourceResp `json:"available_resource,omitempty"`
	HttpStatusCode    int                      `json:"-"`
}

func (o ShowAvailableResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableResourceResponse", string(data)}, " ")
}
