package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowChannelStatisticResponse Response Object
type ShowChannelStatisticResponse struct {

	// 频道推流域名。
	Domain *string `json:"domain,omitempty"`

	// 组名或应用名。
	AppName *string `json:"app_name,omitempty"`

	// 频道ID。频道唯一标识，为必填项。
	Id *string `json:"id,omitempty"`

	// 统计信息的类型，scte35。
	Type *ShowChannelStatisticResponseType `json:"type,omitempty"`

	Scte35         *Scte35StatisticRsp `json:"scte35,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowChannelStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChannelStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowChannelStatisticResponse", string(data)}, " ")
}

type ShowChannelStatisticResponseType struct {
	value string
}

type ShowChannelStatisticResponseTypeEnum struct {
	SCTE35 ShowChannelStatisticResponseType
}

func GetShowChannelStatisticResponseTypeEnum() ShowChannelStatisticResponseTypeEnum {
	return ShowChannelStatisticResponseTypeEnum{
		SCTE35: ShowChannelStatisticResponseType{
			value: "scte35",
		},
	}
}

func (c ShowChannelStatisticResponseType) Value() string {
	return c.value
}

func (c ShowChannelStatisticResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowChannelStatisticResponseType) UnmarshalJSON(b []byte) error {
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
