package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 维护窗口
type CdmQueryClusterDetailsRepsonseMaintainWindow struct {

	// 周几
	Day *string `json:"day,omitempty" xml:"day"`

	// 开始时间。
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 结束时间。
	EndTime *string `json:"endTime,omitempty" xml:"endTime"`
}

func (o CdmQueryClusterDetailsRepsonseMaintainWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmQueryClusterDetailsRepsonseMaintainWindow struct{}"
	}

	return strings.Join([]string{"CdmQueryClusterDetailsRepsonseMaintainWindow", string(data)}, " ")
}
