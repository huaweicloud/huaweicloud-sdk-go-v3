package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceJobsRequest Request Object
type ListResourceJobsRequest struct {

	// - 参数解释：ESW实例的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ResourceId string `json:"resource_id"`

	// - 参数解释：分页查询每页返回的记录个数。 - 约束限制：limit需要和marker配合使用，详细规则请见marker的参数说明。 - 取值范围：0~2^31-1。 - 默认取值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// - 参数解释：分页查询的起始资源ID，表示从指定资源的下一条记录开始查询。 - 约束限制：marker需要和limit配合使用。   - 若不传入marker和limit参数，查询结果返回第一页全部资源记录。   - 若不传入marker参数，limit为10，查询结果返回第1~10条资源记录。   - 若marker为第10条记录的资源ID，limit为10，查询结果返回第11~20条资源记录。   - 若marker为第10条记录的资源ID，不传入limit参数，查询结果返回第11~2000条（limit默认值2000）资源记录。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResourceJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceJobsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceJobsRequest", string(data)}, " ")
}
