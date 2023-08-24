package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAimMsgTemplateResponse Response Object
type DeleteAimMsgTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAimMsgTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAimMsgTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAimMsgTemplateResponse", string(data)}, " ")
}
