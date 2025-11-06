package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PromInstanceRequestModel struct {

	// Prometheus实例名称 名称不能以下划线或中划线开头结尾，只含有中文、英文、数字、下划线、中划线、长度1-100。
	PromName string `json:"prom_name"`

	// Prometheus实例类型。 - ECS：Prometheus for ECS - CCE：Prometheus for CCE - REMOTE_WRITE：Prometheus 通用实例 - CLOUD_SERVICE：Prometheus for 云服务 - ACROSS_ACCOUNT：Prometheus for 多账号聚合实例 [（暂不支持ACROSS_ACCOUNT类型）](tag:hws_eu,g42,sbc,OCB,ctc,cmcc,srg,hk_sbc,ctc,DT)
	PromType string `json:"prom_type"`

	// Prometheus实例版本号。
	PromVersion *string `json:"prom_version,omitempty"`

	// Prometheus实例所属的企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// Prometheus实例所属projectId。
	ProjectId *string `json:"project_id,omitempty"`

	// 被聚合的账号和普罗实例列表。
	AggrPrometheusInfo *[]AggrPrometheusInfo `json:"aggr_prometheus_info,omitempty"`
}

func (o PromInstanceRequestModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PromInstanceRequestModel struct{}"
	}

	return strings.Join([]string{"PromInstanceRequestModel", string(data)}, " ")
}
