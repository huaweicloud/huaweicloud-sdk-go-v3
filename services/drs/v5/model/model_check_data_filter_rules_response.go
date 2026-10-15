package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataFilterRulesResponse Response Object
type CheckDataFilterRulesResponse struct {

	// 异步校验任务的查询标识。提交数据过滤规则校验请求后返回的任务ID，用于后续通过查询/v5/{project_id}/job/{job_id}/data-filtering/result接口获取校验结果。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckDataFilterRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataFilterRulesResponse struct{}"
	}

	return strings.Join([]string{"CheckDataFilterRulesResponse", string(data)}, " ")
}
