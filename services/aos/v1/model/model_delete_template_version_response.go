package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTemplateVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateVersionResponse", string(data)}, " ")
}
