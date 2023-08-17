package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussMysqlDnsRequest Request Object
type CreateGaussMysqlDnsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateDnsNameReq `json:"body,omitempty"`
}

func (o CreateGaussMysqlDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMysqlDnsRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMysqlDnsRequest", string(data)}, " ")
}
