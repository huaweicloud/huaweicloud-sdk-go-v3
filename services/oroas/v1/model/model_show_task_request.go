package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRequest Request Object
type ShowTaskRequest struct {

	// 用户凭证
	XApigAppCode string `json:"X-Apig-AppCode"`

	// 服务类别，针对不同服务类场景，为用户提前填充对应值，用户侧不需单独赋值；当前仅支持 二维切割 2dcut ，便于后续扩展
	ServiceGroup string `json:"service_group"`

	// 子服务类型，针对不同服务，为用户提前填充对应值，用户侧不需单独赋值；服装切割固定为 irregular-textile，雕刻机切割固定为 engraving-machine-cutting， 板材切割固定为 regular-plate
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
