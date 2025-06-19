package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobRequest Request Object
type ShowSparkJobRequest struct {

	// 参数解释:  批处理作业的ID 示例: 0a324461-d9d9-45da-a52a-3b3c7a3d809e 约束限制:  无 取值范围: 无 默认取值: 无
	BatchId string `json:"batch_id"`
}

func (o ShowSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobRequest", string(data)}, " ")
}
