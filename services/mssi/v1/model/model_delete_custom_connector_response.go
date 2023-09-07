package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomConnectorResponse Response Object
type DeleteCustomConnectorResponse struct {

	// 状态码
	ResCode *int32 `json:"res_code,omitempty"`

	// 成功信息
	ResLog *string `json:"res_log,omitempty"`

	// 成功信息
	ResMsg         *string `json:"res_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteCustomConnectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomConnectorResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomConnectorResponse", string(data)}, " ")
}
