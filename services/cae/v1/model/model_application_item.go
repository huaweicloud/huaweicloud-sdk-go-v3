package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationItem 应用信息。
type ApplicationItem struct {

	// ID。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`
}

func (o ApplicationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationItem struct{}"
	}

	return strings.Join([]string{"ApplicationItem", string(data)}, " ")
}
