package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobStatusRequest Request Object
type ShowSqlJobStatusRequest struct {

	// 参数解释:  作业ID 示例: 6d2146a0-c2d5-41bd-8ca0-ca9694ada992 约束限制:  无 取值范围: 无 默认取值: 无
	JobId string `json:"job_id"`
}

func (o ShowSqlJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobStatusRequest", string(data)}, " ")
}
