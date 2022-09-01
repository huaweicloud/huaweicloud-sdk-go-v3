package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMaintenanceWindowsResponse struct {

	// 支持的维护时间窗列表。
	MaintainWindows *[]MaintainWindowsEntity `json:"maintain_windows,omitempty" xml:"maintain_windows"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListMaintenanceWindowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMaintenanceWindowsResponse struct{}"
	}

	return strings.Join([]string{"ListMaintenanceWindowsResponse", string(data)}, " ")
}
