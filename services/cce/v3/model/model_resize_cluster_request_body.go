package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResizeClusterRequestBody 变更集群规格的请求体
type ResizeClusterRequestBody struct {

	// 要变更的目标规格。仅支持变更集群最大节点规模，不支持变更控制节点数，且不支持降低集群规格。例如原集群规格为cce.s2.medium，仅支持变更至cce.s2.large及以上规格，不支持变更至cce.s2.small或cce.s1.medium。  - cce.s1.small: 小规模单控制节点CCE集群（最大50节点） - cce.s1.medium: 中等规模单控制节点CCE集群（最大200节点） - cce.s2.small: 小规模多控制节点CCE集群（最大50节点） - cce.s2.medium: 中等规模多控制节点CCE集群（最大200节点） - cce.s2.large: 大规模多控制节点CCE集群（最大1000节点） - cce.s2.xlarge: 超大规模多控制节点CCE集群（最大2000节点）  >    关于规格参数中的字段说明如下： >    - s1：单控制节点的集群，控制节点数为1。单控制节点故障后，集群将不可用，但已运行工作负载不受影响。 >    - s2：多控制节点的集群，即高可用集群，控制节点数为3。当某个控制节点故障时，集群仍然可用。 >    [- dec：表示专属云的CCE集群规格。例如cce.dec.s1.small表示小规模单控制节点的专属云CCE集群（最大50节点）。](tag:hws,hws_hk) >    - small：表示集群支持管理的最大节点规模为50节点。 >    - medium：表示集群支持管理的最大节点规模为200节点。 >    - large：表示集群支持管理的最大节点规模为1000节点。 >    - xlarge：表示集群支持管理的最大节点规模为2000节点。。
	FlavorResize string `json:"flavorResize"`

	// **参数解释**： 该参数用于控制集群规格变更时跳过部分任务。 **约束限制**： 无 **取值范围**： - IngressChecker: 集群规格变更时跳过Ingress与ELB配置一致性检查  > - 跳过Ingress与ELB配置一致性检查可能导致业务中断，请谨慎操作！ > - 集群不可用或者过载时，必须跳过Ingress与ELB配置一致性检查，否则会导致集群规格变更失败。[请确保变更集群规格前已按 [ELB Ingress与ELB配置不一致如何处理？](https://support.huaweicloud.com/cce_faq/cce_faq_00493.html) 指南解决一致性问题。](tag:hws)[请确保变更集群规格前已按 [ELB Ingress与ELB配置不一致如何处理？](https://support.huaweicloud.com/intl/zh-cn/cce_faq/cce_faq_00493.html) 指南解决一致性问题。](tag:hws_hk)  **默认取值**： 集群不可用时默认包含IngressChecker
	SkippedTasks *[]ResizeClusterRequestBodySkippedTasks `json:"skippedTasks,omitempty"`

	ExtendParam *ResizeClusterRequestBodyExtendParam `json:"extendParam,omitempty"`
}

func (o ResizeClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequestBody", string(data)}, " ")
}

type ResizeClusterRequestBodySkippedTasks struct {
	value string
}

type ResizeClusterRequestBodySkippedTasksEnum struct {
	INGRESS_CHECKER ResizeClusterRequestBodySkippedTasks
}

func GetResizeClusterRequestBodySkippedTasksEnum() ResizeClusterRequestBodySkippedTasksEnum {
	return ResizeClusterRequestBodySkippedTasksEnum{
		INGRESS_CHECKER: ResizeClusterRequestBodySkippedTasks{
			value: "IngressChecker",
		},
	}
}

func (c ResizeClusterRequestBodySkippedTasks) Value() string {
	return c.value
}

func (c ResizeClusterRequestBodySkippedTasks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeClusterRequestBodySkippedTasks) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
