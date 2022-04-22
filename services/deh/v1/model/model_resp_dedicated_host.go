package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 专属主机结构体。
type RespDedicatedHost struct {

	// 专属主机ID。
	DedicatedHostId string `json:"dedicated_host_id"`

	// 专属主机的名称。  长度限制在255个字符以内，并且不能以空格开头或结尾。
	Name string `json:"name"`

	// 在创建云服务器时（未指定专属主机ID），是否允许云服务器自动分配在一台可用的专属主机上。
	AutoPlacement RespDedicatedHostAutoPlacement `json:"auto_placement"`

	// 专属主机所属的可用区。
	AvailabilityZone string `json:"availability_zone"`

	// 专属主机所属的租户ID。
	ProjectId string `json:"project_id"`

	HostProperties *RespHostProperty `json:"host_properties"`

	// 专属主机状态。
	State RespDedicatedHostState `json:"state"`

	// 专属主机可用的vCPU核数。
	AvailableVcpus int32 `json:"available_vcpus"`

	// 专属主机可用的内存大小。
	AvailableMemory int32 `json:"available_memory"`

	// 专属主机的分配时间。
	AllocatedAt string `json:"allocated_at"`

	// 专属主机的释放时间。
	ReleasedAt string `json:"released_at"`

	// 专属主机上的实例总数。
	InstanceTotal int32 `json:"instance_total"`

	// 专属主机上的实例UUID。  查询专属主机列表接口不显示此参数。
	InstanceUuids []string `json:"instance_uuids"`

	// 专属主机标签。
	Tags *interface{} `json:"tags"`

	// 专属主机系统标签。
	SysTags *interface{} `json:"sys_tags"`
}

func (o RespDedicatedHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespDedicatedHost struct{}"
	}

	return strings.Join([]string{"RespDedicatedHost", string(data)}, " ")
}

type RespDedicatedHostAutoPlacement struct {
	value string
}

type RespDedicatedHostAutoPlacementEnum struct {
	ON  RespDedicatedHostAutoPlacement
	OFF RespDedicatedHostAutoPlacement
}

func GetRespDedicatedHostAutoPlacementEnum() RespDedicatedHostAutoPlacementEnum {
	return RespDedicatedHostAutoPlacementEnum{
		ON: RespDedicatedHostAutoPlacement{
			value: "on",
		},
		OFF: RespDedicatedHostAutoPlacement{
			value: "off",
		},
	}
}

func (c RespDedicatedHostAutoPlacement) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespDedicatedHostAutoPlacement) UnmarshalJSON(b []byte) error {
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

type RespDedicatedHostState struct {
	value string
}

type RespDedicatedHostStateEnum struct {
	AVAILABLE RespDedicatedHostState
	RELEASED  RespDedicatedHostState
	FAULT     RespDedicatedHostState
}

func GetRespDedicatedHostStateEnum() RespDedicatedHostStateEnum {
	return RespDedicatedHostStateEnum{
		AVAILABLE: RespDedicatedHostState{
			value: "available",
		},
		RELEASED: RespDedicatedHostState{
			value: "released",
		},
		FAULT: RespDedicatedHostState{
			value: "fault",
		},
	}
}

func (c RespDedicatedHostState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespDedicatedHostState) UnmarshalJSON(b []byte) error {
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
