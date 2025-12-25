package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLtsCloudLogResponse Response Object
type DeleteLtsCloudLogResponse struct {

	// 日志名称
	SourceName *string `json:"source_name,omitempty"`

	// 配置名称
	ConfigName *string `json:"config_name,omitempty"`

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteLtsCloudLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsCloudLogResponse struct{}"
	}

	return strings.Join([]string{"DeleteLtsCloudLogResponse", string(data)}, " ")
}
