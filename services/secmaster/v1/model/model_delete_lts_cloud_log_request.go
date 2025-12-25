package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLtsCloudLogRequest Request Object
type DeleteLtsCloudLogRequest struct {

	// 日志分类
	Csvc string `json:"csvc"`

	// 日志名称
	SourceName string `json:"source_name"`
}

func (o DeleteLtsCloudLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsCloudLogRequest struct{}"
	}

	return strings.Join([]string{"DeleteLtsCloudLogRequest", string(data)}, " ")
}
