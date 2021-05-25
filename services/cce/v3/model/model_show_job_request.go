package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobRequest struct {
	// 作业ID，获取方式请参见[[如何获取接口URI中参数](https://support.huaweicloud.com/api-cce/cce_02_0271.html)](tag:hws)[[如何获取接口URI中参数](https://support.huaweicloud.com/intl/zh-cn/api-cce/cce_02_0271.html)](tag:hws_hk)

	JobId string `json:"job_id"`
}

func (o ShowJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
