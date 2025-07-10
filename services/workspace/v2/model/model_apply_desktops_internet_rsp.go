package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDesktopsInternetRsp 创建云办公带宽响应。
type ApplyDesktopsInternetRsp struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty"`
}

func (o ApplyDesktopsInternetRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDesktopsInternetRsp struct{}"
	}

	return strings.Join([]string{"ApplyDesktopsInternetRsp", string(data)}, " ")
}
