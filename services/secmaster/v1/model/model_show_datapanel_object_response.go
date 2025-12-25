package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatapanelObjectResponse Response Object
type ShowDatapanelObjectResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	Data *DataObjectDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDatapanelObjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatapanelObjectResponse struct{}"
	}

	return strings.Join([]string{"ShowDatapanelObjectResponse", string(data)}, " ")
}
