package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TaskNodeGroups struct {

	// Task节点节点数量，取值范围0～500，Core与Task节点总数最大为500个。
	NodeNum int32 `json:"node_num"`

	// Task节点的实例规格，例如：c3.4xlarge.2.linux.bigdata。实例规格详细说明请参见[MRS所使用的弹性云服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9006.html)和[MRS所使用的裸金属服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9001.html)。 该参数建议从MRS控制台的集群创建页面获取对应区域对应版本所支持的规格。
	NodeSize string `json:"node_size"`

	// Task节点数据磁盘存储类别，目前支持SATA、SAS和SSD。 - SATA：普通IO - SAS：高IO - SSD：超高IO - GPSSD：通用型SSD
	DataVolumeType TaskNodeGroupsDataVolumeType `json:"data_volume_type"`

	// Task节点数据磁盘存储数目，取值范围：0～10。
	DataVolumeCount int32 `json:"data_volume_count"`

	// Task节点数据磁盘存储大小。  取值范围：100GB～32000GB，传值只需填数字，不需要带单位GB。
	DataVolumeSize int32 `json:"data_volume_size"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
}

func (o TaskNodeGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskNodeGroups struct{}"
	}

	return strings.Join([]string{"TaskNodeGroups", string(data)}, " ")
}

type TaskNodeGroupsDataVolumeType struct {
	value string
}

type TaskNodeGroupsDataVolumeTypeEnum struct {
	SATA  TaskNodeGroupsDataVolumeType
	SAS   TaskNodeGroupsDataVolumeType
	SSD   TaskNodeGroupsDataVolumeType
	GPSSD TaskNodeGroupsDataVolumeType
}

func GetTaskNodeGroupsDataVolumeTypeEnum() TaskNodeGroupsDataVolumeTypeEnum {
	return TaskNodeGroupsDataVolumeTypeEnum{
		SATA: TaskNodeGroupsDataVolumeType{
			value: "SATA",
		},
		SAS: TaskNodeGroupsDataVolumeType{
			value: "SAS",
		},
		SSD: TaskNodeGroupsDataVolumeType{
			value: "SSD",
		},
		GPSSD: TaskNodeGroupsDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c TaskNodeGroupsDataVolumeType) Value() string {
	return c.value
}

func (c TaskNodeGroupsDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskNodeGroupsDataVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
