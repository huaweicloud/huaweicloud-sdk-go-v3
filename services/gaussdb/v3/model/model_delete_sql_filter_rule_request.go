package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlFilterRuleRequest Request Object
type DeleteSqlFilterRuleRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteSqlFilterRuleReq `json:"body,omitempty"`
}

func (o DeleteSqlFilterRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlFilterRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlFilterRuleRequest", string(data)}, " ")
}
