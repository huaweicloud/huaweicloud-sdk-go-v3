package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataNodeRelation struct {

	// 源实例id。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 目标实例id。
	TargetInstanceId *string `json:"target_instance_id,omitempty"`
}

func (o DataNodeRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataNodeRelation struct{}"
	}

	return strings.Join([]string{"DataNodeRelation", string(data)}, " ")
}
