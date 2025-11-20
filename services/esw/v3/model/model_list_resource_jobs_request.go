package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceJobsRequest Request Object
type ListResourceJobsRequest struct {

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ResourceId string `json:"resource_id"`
}

func (o ListResourceJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceJobsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceJobsRequest", string(data)}, " ")
}
