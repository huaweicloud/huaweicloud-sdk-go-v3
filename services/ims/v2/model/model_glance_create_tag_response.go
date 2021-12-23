package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceCreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceCreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagResponse", string(data)}, " ")
}
