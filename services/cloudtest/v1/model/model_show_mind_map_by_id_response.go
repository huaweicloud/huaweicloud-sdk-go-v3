package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindMapByIdResponse Response Object
type ShowMindMapByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *MindmapObject `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMindMapByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindMapByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowMindMapByIdResponse", string(data)}, " ")
}
