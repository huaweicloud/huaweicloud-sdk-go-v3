package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreRedisPitrRequestBody struct {

	// 恢复的指定时间点, 格式为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。  获取方法请参见 查询Redis可恢复时间点 中响应“restore_time”字段下参数的值。
	RestoreTime string `json:"restore_time"`
}

func (o RestoreRedisPitrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedisPitrRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreRedisPitrRequestBody", string(data)}, " ")
}
