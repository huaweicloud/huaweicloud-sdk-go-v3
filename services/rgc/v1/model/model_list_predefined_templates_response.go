package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPredefinedTemplatesResponse Response Object
type ListPredefinedTemplatesResponse struct {
	Templates      *[]PredefinedTemplate `json:"templates,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListPredefinedTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPredefinedTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListPredefinedTemplatesResponse", string(data)}, " ")
}
