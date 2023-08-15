package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLabelPageResponse Response Object
type CreateLabelPageResponse struct {

	// 标签页面id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLabelPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelPageResponse struct{}"
	}

	return strings.Join([]string{"CreateLabelPageResponse", string(data)}, " ")
}
