package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlProxyRequest Request Object
type DeleteGaussMySqlProxyRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CloseMysqlProxyRequestBody `json:"body,omitempty"`
}

func (o DeleteGaussMySqlProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlProxyRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlProxyRequest", string(data)}, " ")
}
