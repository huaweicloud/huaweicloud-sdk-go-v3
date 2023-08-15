package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkHoursType struct {

	// 工时类型id
	Id *int32 `json:"id,omitempty"`

	// 工时类型名称
	Name *string `json:"name,omitempty"`

	// 工时类型状态，1表示启用中，2表示未启用
	Status *int32 `json:"status,omitempty"`
}

func (o WorkHoursType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkHoursType struct{}"
	}

	return strings.Join([]string{"WorkHoursType", string(data)}, " ")
}
