package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyGaussMysqlDnsRequest Request Object
type ModifyGaussMysqlDnsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyDnsNameReq `json:"body,omitempty"`
}

func (o ModifyGaussMysqlDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGaussMysqlDnsRequest struct{}"
	}

	return strings.Join([]string{"ModifyGaussMysqlDnsRequest", string(data)}, " ")
}
