package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapRecycleByIdResponse Response Object
type ShowMindmapRecycleByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMindmapRecycleByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapRecycleByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowMindmapRecycleByIdResponse", string(data)}, " ")
}
