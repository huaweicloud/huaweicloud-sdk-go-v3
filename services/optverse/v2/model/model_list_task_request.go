package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskRequest Request Object
type ListTaskRequest struct {

	// 服务类型，针对不同服务，为用户提前填充对应值，用户侧不需单独赋值；二维切割-方形件固定为 regular-plate，二维切割-异形件固定为 irregular-textile， 数学规划求解器固定为 optverse-mpsolver
	ServiceType string `json:"service_type"`

	// 限制量，单次查询总量，必须由数字组成，默认为300，取值范围[1,300]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRequest struct{}"
	}

	return strings.Join([]string{"ListTaskRequest", string(data)}, " ")
}
