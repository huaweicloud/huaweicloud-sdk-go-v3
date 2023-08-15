package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateResponse Response Object
type UpdateTemplateResponse struct {

	// 模板ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateResponse", string(data)}, " ")
}
