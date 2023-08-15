package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceDeleteImageMemberResponse Response Object
type GlanceDeleteImageMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageMemberResponse", string(data)}, " ")
}
