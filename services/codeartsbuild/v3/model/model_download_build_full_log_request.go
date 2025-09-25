package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBuildFullLogRequest Request Object
type DownloadBuildFullLogRequest struct {

	// 记录ID,36位数字、小写字母、'-'组组合。
	RecordId string `json:"record_id"`

	// 日志等级 值为INFO | DEBUG。
	LogLevel *string `json:"log_level,omitempty"`

	// **参数解释**： 是否压缩日志。 **约束限制**： 不涉及。 **取值范围**： true|false。 **默认取值**： false
	Compress *bool `json:"compress,omitempty"`
}

func (o DownloadBuildFullLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBuildFullLogRequest struct{}"
	}

	return strings.Join([]string{"DownloadBuildFullLogRequest", string(data)}, " ")
}
