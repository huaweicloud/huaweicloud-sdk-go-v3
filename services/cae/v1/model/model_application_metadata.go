package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据。
type ApplicationMetadata struct {

	// 应用id。
	Id *string `json:"id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 附加应用信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ApplicationMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationMetadata struct{}"
	}

	return strings.Join([]string{"ApplicationMetadata", string(data)}, " ")
}
