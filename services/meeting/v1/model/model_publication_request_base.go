package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新增/更新发布信息请求
type PublicationRequestBase struct {

	// 发布名称
	PublishName string `json:"publishName" xml:"publishName"`

	// 开始时间
	StartTime int64 `json:"startTime" xml:"startTime"`

	// 结束时间
	EndTime int64 `json:"endTime" xml:"endTime"`

	// 发布到部门编码列表
	DeptList []string `json:"deptList" xml:"deptList"`

	// 发布到设备用户ID列表
	DeviceList []string `json:"deviceList" xml:"deviceList"`
}

func (o PublicationRequestBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationRequestBase struct{}"
	}

	return strings.Join([]string{"PublicationRequestBase", string(data)}, " ")
}
