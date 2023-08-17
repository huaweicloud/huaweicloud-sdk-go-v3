package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDataOptionalParam 添加数据的可选参数，其中每个参数仅对部分服务类型生效，具体可参见服务类型说明。
type AddDataOptionalParam struct {

	// 是否进行目标检测，默认为true。
	DoDet *bool `json:"do_det,omitempty"`

	// 目标矩形框坐标，如给定则不进行目标检测，直接使用该box作为目标。格式为“x1,y1,x2,y2”（无空格），x1/y1为目标左上角坐标，x2/y2为目标右下角坐标，具体要求如下： - 0 <= x1 < x2 <= width，默认要求x2-x1 >= 15，具体可参考服务类型说明。 - 0 <= y1 < y2 <= height，默认要求y2-y1 >= 15，具体可参考服务类型说明。
	Box *string `json:"box,omitempty"`

	// 是否进行对象分类，默认为true。
	DoCls *bool `json:"do_cls,omitempty"`

	// 对象类目，如给定则不进行对象分类，直接使用该category作为类目。具体类目信息可参见对应的服务类型说明。
	Category float32 `json:"category,omitempty"`
}

func (o AddDataOptionalParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataOptionalParam struct{}"
	}

	return strings.Join([]string{"AddDataOptionalParam", string(data)}, " ")
}
