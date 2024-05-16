package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapStorageTypeRequest Request Object
type ListHtapStorageTypeRequest struct {

	// HTAP数据库名。 取值范围： - star-rocks - click-house
	Database string `json:"database"`

	// 数据库大版本号
	VersionName string `json:"version_name"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListHtapStorageTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapStorageTypeRequest struct{}"
	}

	return strings.Join([]string{"ListHtapStorageTypeRequest", string(data)}, " ")
}
