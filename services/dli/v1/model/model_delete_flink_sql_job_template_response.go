package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkSqlJobTemplateResponse Response Object
type DeleteFlinkSqlJobTemplateResponse struct {

	// 响应正确与否的标志，true表示成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	Template       *FlinkSqlJobTemplateId `json:"template,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o DeleteFlinkSqlJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkSqlJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlinkSqlJobTemplateResponse", string(data)}, " ")
}
