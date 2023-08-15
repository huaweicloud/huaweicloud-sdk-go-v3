package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateGroupCollectionRequest Request Object
type CreateTemplateGroupCollectionRequest struct {
	Body *TransTemplateGroupCollection `json:"body,omitempty"`
}

func (o CreateTemplateGroupCollectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupCollectionRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupCollectionRequest", string(data)}, " ")
}
