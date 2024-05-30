package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStandardTemplateResponse Response Object
type DeleteStandardTemplateResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteStandardTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStandardTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteStandardTemplateResponse", string(data)}, " ")
}
