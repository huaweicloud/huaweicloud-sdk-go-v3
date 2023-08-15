package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearGraph2Request Request Object
type ClearGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 是否清空图关联的元数据。建议清除。
	ClearMetadata *bool `json:"clear_metadata,omitempty"`
}

func (o ClearGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearGraph2Request struct{}"
	}

	return strings.Join([]string{"ClearGraph2Request", string(data)}, " ")
}
