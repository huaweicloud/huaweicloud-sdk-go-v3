package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AggrPrometheusInfo struct {

	// 被聚合的账号的projectid。
	ProjectId string `json:"project_id"`

	// 被聚合的账号下的普罗ID列表。
	PrometheusIds []string `json:"prometheusIds"`

	// 被聚合的账号id。
	Id string `json:"id"`

	// 被聚合的账号名称。
	Name string `json:"name"`
}

func (o AggrPrometheusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggrPrometheusInfo struct{}"
	}

	return strings.Join([]string{"AggrPrometheusInfo", string(data)}, " ")
}
