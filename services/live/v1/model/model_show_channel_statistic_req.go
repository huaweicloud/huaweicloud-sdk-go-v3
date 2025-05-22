package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowChannelStatisticReq 查询OTT频道统计信息消息体。
type ShowChannelStatisticReq struct {

	// 频道推流域名。
	Domain string `json:"domain"`

	// 组名或应用名。
	AppName string `json:"app_name"`

	// 频道ID。频道唯一标识，为必填项。
	Id string `json:"id"`

	// 统计信息的类型，scte35。
	Type ShowChannelStatisticReqType `json:"type"`

	Scte35 *Scte35StatisticReq `json:"scte35,omitempty"`
}

func (o ShowChannelStatisticReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChannelStatisticReq struct{}"
	}

	return strings.Join([]string{"ShowChannelStatisticReq", string(data)}, " ")
}

type ShowChannelStatisticReqType struct {
	value string
}

type ShowChannelStatisticReqTypeEnum struct {
	SCTE35 ShowChannelStatisticReqType
}

func GetShowChannelStatisticReqTypeEnum() ShowChannelStatisticReqTypeEnum {
	return ShowChannelStatisticReqTypeEnum{
		SCTE35: ShowChannelStatisticReqType{
			value: "scte35",
		},
	}
}

func (c ShowChannelStatisticReqType) Value() string {
	return c.value
}

func (c ShowChannelStatisticReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowChannelStatisticReqType) UnmarshalJSON(b []byte) error {
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
