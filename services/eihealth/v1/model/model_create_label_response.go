package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLabelResponse Response Object
type CreateLabelResponse struct {

	// 标签id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelResponse struct{}"
	}

	return strings.Join([]string{"CreateLabelResponse", string(data)}, " ")
}
