package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAccessConfigResponse Response Object
type CreateAccessConfigResponse struct {

	// 日志接入ID
	AccessConfigId *string `json:"access_config_id,omitempty"`

	// 日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty"`

	// 日志接入类型。AGENT：ECS接入  K8S_CCE: CCE接入
	AccessConfigType *CreateAccessConfigResponseAccessConfigType `json:"access_config_type,omitempty"`

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

	// CCE集群ID
	ClusterId      *string `json:"cluster_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAccessConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessConfigResponse", string(data)}, " ")
}

type CreateAccessConfigResponseAccessConfigType struct {
	value string
}

type CreateAccessConfigResponseAccessConfigTypeEnum struct {
	AGENT    CreateAccessConfigResponseAccessConfigType
	K8_S_CCE CreateAccessConfigResponseAccessConfigType
}

func GetCreateAccessConfigResponseAccessConfigTypeEnum() CreateAccessConfigResponseAccessConfigTypeEnum {
	return CreateAccessConfigResponseAccessConfigTypeEnum{
		AGENT: CreateAccessConfigResponseAccessConfigType{
			value: "AGENT",
		},
		K8_S_CCE: CreateAccessConfigResponseAccessConfigType{
			value: "K8S_CCE",
		},
	}
}

func (c CreateAccessConfigResponseAccessConfigType) Value() string {
	return c.value
}

func (c CreateAccessConfigResponseAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessConfigResponseAccessConfigType) UnmarshalJSON(b []byte) error {
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
