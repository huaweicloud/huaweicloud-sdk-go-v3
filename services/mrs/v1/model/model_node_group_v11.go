package model

import (
	"encoding/json"

	"strings"
)

type NodeGroupV11 struct {
	// 节点组名。 - master_node_default_group - core_node_analysis_group - core_node_streaming_group - task_node_analysis_group - task_node_streaming_group

	GroupName string `json:"group_name"`
	// 节点数量，取值范围0～500，Core与Task节点总数最大为500个。

	NodeNum int32 `json:"node_num"`
	// 节点的实例规格。配置方法请参考“master_node_size”配置的备注。

	NodeSize string `json:"node_size"`
	// 节点系统磁盘存储大小。

	RootVolumeSize *string `json:"root_volume_size,omitempty"`
	// 节点系统磁盘存储类别，目前支持SATA、SAS和SSD。 - SATA：普通IO - SAS：高IO - SSD：超高IO - GPSSD：通用型SSD

	RootVolumeType *string `json:"root_volume_type,omitempty"`
	// 节点数据磁盘存储类别，目前支持SATA、SAS和SSD。 - SATA：普通IO - SAS：高IO - SSD：超高IO - GPSSD：通用型SSD

	DataVolumeType *string `json:"data_volume_type,omitempty"`
	// 节点数据磁盘存储数目 取值范围：0～10。

	DataVolumeCount *int32 `json:"data_volume_count,omitempty"`
	// 节点数据磁盘存储大小 取值范围：100GB～32000GB。

	DataVolumeSize *int32 `json:"data_volume_size,omitempty"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
}

func (o NodeGroupV11) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeGroupV11 struct{}"
	}

	return strings.Join([]string{"NodeGroupV11", string(data)}, " ")
}
