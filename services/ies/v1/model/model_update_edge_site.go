package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEdgeSite struct {

	// 边缘小站描述，最大支持长度为255个字节，不允许包含<>
	Description *string `json:"description,omitempty"`

	Location *UpdateLocation `json:"location,omitempty"`
}

func (o UpdateEdgeSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeSite struct{}"
	}

	return strings.Join([]string{"UpdateEdgeSite", string(data)}, " ")
}
