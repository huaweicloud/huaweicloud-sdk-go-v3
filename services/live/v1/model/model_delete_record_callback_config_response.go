package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordCallbackConfigResponse Response Object
type DeleteRecordCallbackConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRecordCallbackConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordCallbackConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordCallbackConfigResponse", string(data)}, " ")
}
