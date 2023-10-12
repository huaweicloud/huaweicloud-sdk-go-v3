package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationMetadata 响应数据。
type ApplicationMetadata struct {

	// 应用ID。
	Id *string `json:"id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ApplicationMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationMetadata struct{}"
	}

	return strings.Join([]string{"ApplicationMetadata", string(data)}, " ")
}
