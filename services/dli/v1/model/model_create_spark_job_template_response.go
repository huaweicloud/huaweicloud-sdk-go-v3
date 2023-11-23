package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSparkJobTemplateResponse Response Object
type CreateSparkJobTemplateResponse struct {

	// 模板ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSparkJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateSparkJobTemplateResponse", string(data)}, " ")
}
