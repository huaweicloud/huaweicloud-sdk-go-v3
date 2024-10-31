package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMindmapRecycleByIdResponse Response Object
type DeleteMindmapRecycleByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMindmapRecycleByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMindmapRecycleByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteMindmapRecycleByIdResponse", string(data)}, " ")
}
