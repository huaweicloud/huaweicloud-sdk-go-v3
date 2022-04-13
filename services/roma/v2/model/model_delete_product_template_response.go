package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteProductTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProductTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteProductTemplateResponse", string(data)}, " ")
}
