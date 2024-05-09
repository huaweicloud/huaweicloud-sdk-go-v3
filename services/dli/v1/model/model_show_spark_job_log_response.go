package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobLogResponse Response Object
type ShowSparkJobLogResponse struct {

	// 批处理作业的id。
	Id *string `json:"id,omitempty"`

	// 日志起始索引。
	From *int32 `json:"from,omitempty"`

	// 日志的总记录数。
	Total *int64 `json:"total,omitempty"`

	// 显示当前批处理作业日志。
	Log            *[]string `json:"log,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowSparkJobLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobLogResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobLogResponse", string(data)}, " ")
}
