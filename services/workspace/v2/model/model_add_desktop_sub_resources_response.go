package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopSubResourcesResponse Response Object
type AddDesktopSubResourcesResponse struct {

	// 创建云桌面总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDesktopSubResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopSubResourcesResponse struct{}"
	}

	return strings.Join([]string{"AddDesktopSubResourcesResponse", string(data)}, " ")
}
