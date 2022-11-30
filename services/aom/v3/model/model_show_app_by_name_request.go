package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAppByNameRequest struct {

	// 应用唯一标识
	Name *string `json:"name,omitempty"`

	// 实体的显示名称。
	DisplayName *string `json:"display_name,omitempty"`
}

func (o ShowAppByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppByNameRequest struct{}"
	}

	return strings.Join([]string{"ShowAppByNameRequest", string(data)}, " ")
}
