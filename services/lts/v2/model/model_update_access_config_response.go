package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAccessConfigResponse Response Object
type UpdateAccessConfigResponse struct {

	// 日志接入ID
	AccessConfigId *string `json:"access_config_id,omitempty"`

	// 日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty"`

	// 日志接入类型。AGENT：ECS接入  K8S_CCE: CCE接入
	AccessConfigType *UpdateAccessConfigResponseAccessConfigType `json:"access_config_type,omitempty"`

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
	BinaryCollect  *bool `json:"binary_collect,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateAccessConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigResponse", string(data)}, " ")
}

type UpdateAccessConfigResponseAccessConfigType struct {
	value string
}

type UpdateAccessConfigResponseAccessConfigTypeEnum struct {
	AGENT    UpdateAccessConfigResponseAccessConfigType
	K8_S_CCE UpdateAccessConfigResponseAccessConfigType
}

func GetUpdateAccessConfigResponseAccessConfigTypeEnum() UpdateAccessConfigResponseAccessConfigTypeEnum {
	return UpdateAccessConfigResponseAccessConfigTypeEnum{
		AGENT: UpdateAccessConfigResponseAccessConfigType{
			value: "AGENT",
		},
		K8_S_CCE: UpdateAccessConfigResponseAccessConfigType{
			value: "K8S_CCE",
		},
	}
}

func (c UpdateAccessConfigResponseAccessConfigType) Value() string {
	return c.value
}

func (c UpdateAccessConfigResponseAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAccessConfigResponseAccessConfigType) UnmarshalJSON(b []byte) error {
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
