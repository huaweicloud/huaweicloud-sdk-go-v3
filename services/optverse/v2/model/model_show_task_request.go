package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRequest Request Object
type ShowTaskRequest struct {

	// 服务类型，针对不同服务，为用户提前填充对应值，用户侧不需单独赋值；二维切割-方形件固定为 regular-plate，二维切割-异形件固定为 irregular-textile， 数学规划求解器固定为 optverse-mpsolver
	ServiceType string `json:"service_type"`

	// 任务id
	TaskId string `json:"task_id"`

	// 是否返回输入信息，默认为false； 当前暂不支持该功能，后续扩展
	InputEnable *bool `json:"input_enable,omitempty"`
}

func (o ShowTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRequest", string(data)}, " ")
}
