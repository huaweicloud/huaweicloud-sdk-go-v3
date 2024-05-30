package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesResultData data，统一的返回结果的最外层数据结构。
type ListWorkspacesResultData struct {
	Value *ListWorkspacesResultDataValue `json:"value,omitempty"`
}

func (o ListWorkspacesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesResultData struct{}"
	}

	return strings.Join([]string{"ListWorkspacesResultData", string(data)}, " ")
}
