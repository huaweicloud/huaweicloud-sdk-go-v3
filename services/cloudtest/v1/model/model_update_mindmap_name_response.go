package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMindmapNameResponse Response Object
type UpdateMindmapNameResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMindmapNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMindmapNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateMindmapNameResponse", string(data)}, " ")
}
