package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmRuleTemplateSpecWithCloudService struct {

	// 关联的云服务。
	RelatedCloudService *string `json:"related_cloud_service,omitempty"`

	// 关联的CCE集群。
	RelatedCceClusters *[]string `json:"related_cce_clusters,omitempty"`

	// 关联的Prom实例。
	RelatedPrometheusInstances *[]string `json:"related_prometheus_instances,omitempty"`

	AlarmNotification *AlarmNotification `json:"alarm_notification,omitempty"`

	// 告警模板列表。
	AlarmTemplateSpecItems *[]AlarmTemplateSpecItem `json:"alarm_template_spec_items,omitempty"`
}

func (o AlarmRuleTemplateSpecWithCloudService) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleTemplateSpecWithCloudService struct{}"
	}

	return strings.Join([]string{"AlarmRuleTemplateSpecWithCloudService", string(data)}, " ")
}
