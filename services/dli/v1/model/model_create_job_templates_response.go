package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobTemplatesResponse Response Object
type CreateJobTemplatesResponse struct {

	// 模板ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateJobTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobTemplatesResponse struct{}"
	}

	return strings.Join([]string{"CreateJobTemplatesResponse", string(data)}, " ")
}
