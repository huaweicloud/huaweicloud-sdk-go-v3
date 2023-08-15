package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateGroupCollectionResponse Response Object
type DeleteTemplateGroupCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateGroupCollectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupCollectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupCollectionResponse", string(data)}, " ")
}
