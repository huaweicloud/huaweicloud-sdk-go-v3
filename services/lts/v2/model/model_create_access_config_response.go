package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateAccessConfigResponse struct {
	// 日志接入ID

	AccessConfigId *string `json:"access_config_id,omitempty"`
	// 日志接入名称

	AccessConfigName *string `json:"access_config_name,omitempty"`
	// 日志接入类型。AGENT：主机接入

	AccessConfigType *CreateAccessConfigResponseAccessConfigType `json:"access_config_type,omitempty"`
	// 创建时间

	CreateTime *int64 `json:"create_time,omitempty"`

	AccessConfigDetail *AccessConfigDeatil `json:"access_config_detail,omitempty"`

	LogInfo *AccessConfigQueryLogInfo `json:"log_info,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`
	HttpStatusCode  int                `json:"-"`
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
	AGENT CreateAccessConfigResponseAccessConfigType
}

func GetCreateAccessConfigResponseAccessConfigTypeEnum() CreateAccessConfigResponseAccessConfigTypeEnum {
	return CreateAccessConfigResponseAccessConfigTypeEnum{
		AGENT: CreateAccessConfigResponseAccessConfigType{
			value: "AGENT",
		},
	}
}

func (c CreateAccessConfigResponseAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessConfigResponseAccessConfigType) UnmarshalJSON(b []byte) error {
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
