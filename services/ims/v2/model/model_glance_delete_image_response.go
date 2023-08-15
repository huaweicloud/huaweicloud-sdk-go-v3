package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceDeleteImageResponse Response Object
type GlanceDeleteImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageResponse", string(data)}, " ")
}
