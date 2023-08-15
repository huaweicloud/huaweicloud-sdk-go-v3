package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ClusterScalingParams struct {

	// 扩容/缩容时系统获取的订单号，用户不需要配置。
	OrderId *string `json:"order_id,omitempty"`

	// - scale_in：缩容 - scale_out：扩容
	ScaleType ClusterScalingParamsScaleType `json:"scale_type"`

	// 扩容/缩容时新增或者减少节点的ID标识,参数值固定为node_orderadd。
	NodeId string `json:"node_id"`

	// 扩容或缩容的节点组。 - 如果node_group为core_node_default_group，表示Core节点组。 - 如果node_group为task_node_default_group，表示Task节点组。  该字段可以为空，为空时，系统默认值为core_node_default_group。
	NodeGroup *string `json:"node_group,omitempty"`

	// 是否跳过引导操作，默认为false，即执行引导操作。 仅在创建集群时配置了引导操作且扩容时有意义，表示扩容时是否在新增节点上执行创建集群时指定的引导操作。
	SkipBootstrapScripts *string `json:"skip_bootstrap_scripts,omitempty"`

	// 扩容后是否启动扩容节点上的组件。  - true：扩容后不启动组件。 - false：扩容后启动组件。
	ScaleWithoutStart *bool `json:"scale_without_start,omitempty"`

	// 缩容Task节点时指定待删除Task节点的ID列表。  - 当scale_type为扩容时，该参数不生效。 - 当scale_type为缩容且该参数不为空时，删除指定的Task节点。 - 当scale_type为缩容且server_ids为空时，按照系统规则自动选择删除Task节点。
	ServerIds *[]string `json:"server_ids,omitempty"`

	// 扩容或缩容的节点数。  - 扩容时的最大节点数为（500 - 集群Core/Task节点数）。例如，当前集群Core节点数为3，此处扩容的节点数必须小于等于497。     Core和Task节点总数最大值为500，如果用户需要的Core/Task节点数大于500，可以联系技术支持人员或者调用后台接口修改数据库。   - 缩容时Core节点数大于3或者Task节点数大于0可以进行节点删除。例如，当前集群Core节点和Task节点数均为5，Core节点可缩容的节点数为2（5减去3），Task节点可缩容节点数为小于等于5。
	Instances int32 `json:"instances"`

	TaskNodeInfo *TaskNodeInfo `json:"task_node_info,omitempty"`
}

func (o ClusterScalingParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterScalingParams struct{}"
	}

	return strings.Join([]string{"ClusterScalingParams", string(data)}, " ")
}

type ClusterScalingParamsScaleType struct {
	value string
}

type ClusterScalingParamsScaleTypeEnum struct {
	SCALE_IN  ClusterScalingParamsScaleType
	SCALE_OUT ClusterScalingParamsScaleType
}

func GetClusterScalingParamsScaleTypeEnum() ClusterScalingParamsScaleTypeEnum {
	return ClusterScalingParamsScaleTypeEnum{
		SCALE_IN: ClusterScalingParamsScaleType{
			value: "scale_in",
		},
		SCALE_OUT: ClusterScalingParamsScaleType{
			value: "scale_out",
		},
	}
}

func (c ClusterScalingParamsScaleType) Value() string {
	return c.value
}

func (c ClusterScalingParamsScaleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterScalingParamsScaleType) UnmarshalJSON(b []byte) error {
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
