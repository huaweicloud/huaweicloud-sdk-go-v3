package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateAccessConfigResponse struct {

	// 日志接入ID
	AccessConfigId *string `json:"access_config_id,omitempty" xml:"access_config_id"`

	// 日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty" xml:"access_config_name"`

	// 日志接入类型。AGENT：主机接入
	AccessConfigType *UpdateAccessConfigResponseAccessConfigType `json:"access_config_type,omitempty" xml:"access_config_type"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	AccessConfigDetail *AccessConfigDeatil `json:"access_config_detail,omitempty" xml:"access_config_detail"`

	LogInfo *AccessConfigQueryLogInfo `json:"log_info,omitempty" xml:"log_info"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty" xml:"host_group_info"`

	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty" xml:"access_config_tag"`
	HttpStatusCode  int                `json:"-"`
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
	AGENT UpdateAccessConfigResponseAccessConfigType
}

func GetUpdateAccessConfigResponseAccessConfigTypeEnum() UpdateAccessConfigResponseAccessConfigTypeEnum {
	return UpdateAccessConfigResponseAccessConfigTypeEnum{
		AGENT: UpdateAccessConfigResponseAccessConfigType{
			value: "AGENT",
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
