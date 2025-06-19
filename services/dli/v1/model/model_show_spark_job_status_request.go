package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobStatusRequest Request Object
type ShowSparkJobStatusRequest struct {

	// 参数解释:  批处理作业的ID 示例: 0a324461-d9d9-45da-a52a-3b3c7a3d809e 约束限制:  匹配正则表达式'^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$'的字符串 取值范围: 无 默认取值: 无
	BatchId string `json:"batch_id"`
}

func (o ShowSparkJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobStatusRequest", string(data)}, " ")
}
