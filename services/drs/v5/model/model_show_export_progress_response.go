package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExportProgressResponse Response Object
type ShowExportProgressResponse struct {

	// 导出进度。
	ProcessStatus  *string `json:"process_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExportProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExportProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowExportProgressResponse", string(data)}, " ")
}
