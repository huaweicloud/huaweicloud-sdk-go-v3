package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTemplateMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTemplateMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateMetadataResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateMetadataResponse", string(data)}, " ")
}
