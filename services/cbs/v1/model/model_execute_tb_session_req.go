package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ExecuteTbSessionReq struct {

	// 客户回复。
	Reply *string `json:"reply,omitempty"`

	// 客户回复属性，0表示通用回复，1表示客户打断， 2表示客户长时未回复。
	Type int32 `json:"type"`
}

func (o ExecuteTbSessionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTbSessionReq struct{}"
	}

	return strings.Join([]string{"ExecuteTbSessionReq", string(data)}, " ")
}
