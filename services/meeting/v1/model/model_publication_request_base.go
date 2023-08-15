package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicationRequestBase 发布信息。
type PublicationRequestBase struct {

	// 发布名称。
	PublishName string `json:"publishName"`

	// 开始时间。
	StartTime int64 `json:"startTime"`

	// 结束时间。
	EndTime int64 `json:"endTime"`

	// 发布到部门编码列表。
	DeptList []string `json:"deptList"`

	// 发布到设备用户ID列表。
	DeviceList []string `json:"deviceList"`
}

func (o PublicationRequestBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationRequestBase struct{}"
	}

	return strings.Join([]string{"PublicationRequestBase", string(data)}, " ")
}
