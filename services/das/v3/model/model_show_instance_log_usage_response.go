package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceLogUsageResponse Response Object
type ShowInstanceLogUsageResponse struct {

	// 日志占用量(GB),保留两位小数
	Size *string `json:"size,omitempty"`

	// 计算用量截止时间
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceLogUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceLogUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceLogUsageResponse", string(data)}, " ")
}
