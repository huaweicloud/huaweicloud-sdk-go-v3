package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigInfo 日志接入配置
type AccessConfigInfo struct {

	// 日志接入ID
	AccessConfigId *string `json:"access_config_id,omitempty"`

	// 日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty"`

	// 日志接入类型。AGENT：ECS接入  K8S_CCE: CCE接入
	AccessConfigType *AccessConfigInfoAccessConfigType `json:"access_config_type,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	AccessConfigDetail *AccessConfigDeatilCreate `json:"access_config_detail,omitempty"`

	LogInfo *AccessConfigQueryLogInfo `json:"log_info,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 二进制采集
	LogSplit *bool `json:"log_split,omitempty"`

	// 日志拆分
	BinaryCollect *bool `json:"binary_collect,omitempty"`
}

func (o AccessConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigInfo", string(data)}, " ")
}

type AccessConfigInfoAccessConfigType struct {
	value string
}

type AccessConfigInfoAccessConfigTypeEnum struct {
	AGENT    AccessConfigInfoAccessConfigType
	K8_S_CCE AccessConfigInfoAccessConfigType
}

func GetAccessConfigInfoAccessConfigTypeEnum() AccessConfigInfoAccessConfigTypeEnum {
	return AccessConfigInfoAccessConfigTypeEnum{
		AGENT: AccessConfigInfoAccessConfigType{
			value: "AGENT",
		},
		K8_S_CCE: AccessConfigInfoAccessConfigType{
			value: "K8S_CCE",
		},
	}
}

func (c AccessConfigInfoAccessConfigType) Value() string {
	return c.value
}

func (c AccessConfigInfoAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigInfoAccessConfigType) UnmarshalJSON(b []byte) error {
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
