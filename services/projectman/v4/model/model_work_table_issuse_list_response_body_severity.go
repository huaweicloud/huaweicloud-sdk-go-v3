package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重要程度信息
type WorkTableIssuseListResponseBodySeverity struct {

	// 重要程度id
	Id *int32 `json:"id,omitempty"`

	// 重要程度名称
	Name *string `json:"name,omitempty"`
}

func (o WorkTableIssuseListResponseBodySeverity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodySeverity struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodySeverity", string(data)}, " ")
}
