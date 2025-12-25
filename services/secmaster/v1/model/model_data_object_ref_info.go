package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataObjectRefInfo Dataobject实体信息
type DataObjectRefInfo struct {

	// 流程实例上下文内容
	Content map[string]interface{} `json:"content,omitempty"`

	Dataclass *DataClassRef `json:"dataclass,omitempty"`
}

func (o DataObjectRefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectRefInfo struct{}"
	}

	return strings.Join([]string{"DataObjectRefInfo", string(data)}, " ")
}
