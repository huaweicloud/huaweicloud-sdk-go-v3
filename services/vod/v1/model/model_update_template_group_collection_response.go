package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateGroupCollectionResponse Response Object
type UpdateTemplateGroupCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTemplateGroupCollectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupCollectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupCollectionResponse", string(data)}, " ")
}
