package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSourceRequest Request Object
type DeleteSourceRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 规则ID
	RuleId string `json:"rule_id"`

	// 源数据源ID
	SourceId int32 `json:"source_id"`
}

func (o DeleteSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteSourceRequest", string(data)}, " ")
}
