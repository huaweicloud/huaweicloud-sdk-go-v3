package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProductTemplateResponse Response Object
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
