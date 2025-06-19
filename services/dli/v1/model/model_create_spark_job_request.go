package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSparkJobRequest Request Object
type CreateSparkJobRequest struct {

	// 参数解释:   用户ID 示例: 48cc2c48765f481480c7db940d6409d1 约束限制:  无 取值范围: 无 默认取值: 无
	UserId *string `json:"USER-ID,omitempty"`

	Body *CreateSparkJobRequestBody `json:"body,omitempty"`
}

func (o CreateSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSparkJobRequest", string(data)}, " ")
}
