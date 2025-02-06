package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequest Request Object
type CreateTaskRequest struct {

	// 服务类型，针对不同服务，为用户提前填充对应值，用户侧不需单独赋值；二维切割-方形件固定为 regular-plate，二维切割-异形件固定为 irregular-textile， 数学规划求解器固定为 optverse-mpsolver
	ServiceType string `json:"service_type"`

	Body *TaskDto `json:"body,omitempty"`
}

func (o CreateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskRequest", string(data)}, " ")
}
