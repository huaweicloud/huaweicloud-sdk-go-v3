package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新发布信息请求
type UpdatePublicationRequestDto struct {
	// 发布名称

	PublishName string `json:"publishName"`
	// 开始时间

	StartTime int64 `json:"startTime"`
	// 结束时间

	EndTime int64 `json:"endTime"`
	// 发布到部门编码列表

	DeptList []string `json:"deptList"`
	// 发布到设备用户ID列表

	DeviceList []string `json:"deviceList"`
	// 发布节目ID列表

	ProgramList *[]string `json:"programList,omitempty"`
}

func (o UpdatePublicationRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicationRequestDto struct{}"
	}

	return strings.Join([]string{"UpdatePublicationRequestDto", string(data)}, " ")
}
