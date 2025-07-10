package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopSubResourcesResponse Response Object
type DeleteDesktopSubResourcesResponse struct {

	// 创建云桌面总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDesktopSubResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopSubResourcesResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesktopSubResourcesResponse", string(data)}, " ")
}
