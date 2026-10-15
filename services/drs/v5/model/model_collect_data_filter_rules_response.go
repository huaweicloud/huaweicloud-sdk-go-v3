package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectDataFilterRulesResponse Response Object
type CollectDataFilterRulesResponse struct {

	// 异步查询数据过滤规则的任务ID。需使用此ID轮询获取数据校验规则结果，参考接口/v5/{project_id}/jobs/{job_id}/filter-rules/result。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectDataFilterRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectDataFilterRulesResponse struct{}"
	}

	return strings.Join([]string{"CollectDataFilterRulesResponse", string(data)}, " ")
}
