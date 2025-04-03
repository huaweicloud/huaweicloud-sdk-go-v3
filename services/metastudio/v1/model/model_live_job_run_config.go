package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LiveJobRunConfig 数字人直播任务运行配置
type LiveJobRunConfig struct {

	// 允许使用资源类型。 * PERIOD：使用包周期资源 * ONDEMAND：使用按需资源 * UNLIMITED：不限制资源类型
	AllowResourceType *LiveJobRunConfigAllowResourceType `json:"allow_resource_type,omitempty"`

	// 一个直播间是否仅允许一个正在直播的任务。 * true: 限制直播间仅允许一个任务运行。 * false: 不限制直播间任务运行数量。
	SingleJobInRoom *bool `json:"single_job_in_room,omitempty"`
}

func (o LiveJobRunConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveJobRunConfig struct{}"
	}

	return strings.Join([]string{"LiveJobRunConfig", string(data)}, " ")
}

type LiveJobRunConfigAllowResourceType struct {
	value string
}

type LiveJobRunConfigAllowResourceTypeEnum struct {
	PERIOD    LiveJobRunConfigAllowResourceType
	ONDEMAND  LiveJobRunConfigAllowResourceType
	UNLIMITED LiveJobRunConfigAllowResourceType
}

func GetLiveJobRunConfigAllowResourceTypeEnum() LiveJobRunConfigAllowResourceTypeEnum {
	return LiveJobRunConfigAllowResourceTypeEnum{
		PERIOD: LiveJobRunConfigAllowResourceType{
			value: "PERIOD",
		},
		ONDEMAND: LiveJobRunConfigAllowResourceType{
			value: "ONDEMAND",
		},
		UNLIMITED: LiveJobRunConfigAllowResourceType{
			value: "UNLIMITED",
		},
	}
}

func (c LiveJobRunConfigAllowResourceType) Value() string {
	return c.value
}

func (c LiveJobRunConfigAllowResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveJobRunConfigAllowResourceType) UnmarshalJSON(b []byte) error {
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
