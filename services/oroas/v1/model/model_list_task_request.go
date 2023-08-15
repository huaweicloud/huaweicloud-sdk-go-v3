package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskRequest Request Object
type ListTaskRequest struct {

	// 用户凭证
	XApigAppCode string `json:"X-Apig-AppCode"`

	// 服务类别，针对不同服务类场景，为用户提前填充对应值，用户侧不需单独赋值；当前仅支持 二维切割 2dcut ，便于后续扩展
	ServiceGroup string `json:"service_group"`

	// 子服务类型，针对不同服务，为用户提前填充对应值，用户侧不需单独赋值；服装切割固定为 irregular-textile，雕刻机切割固定为 engraving-machine-cutting， 板材切割固定为 regular-plate
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
