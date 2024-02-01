package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GlobalConnectionBandwidthQuotas struct {

	// 配额大小。
	Quota *int32 `json:"quota,omitempty"`

	// 已使用值。
	Used *int32 `json:"used,omitempty"`

	// 配额类型。 gcb.size：全域互联带宽大小 gcb.count：全域互联带宽数量
	Type *GlobalConnectionBandwidthQuotasType `json:"type,omitempty"`
}

func (o GlobalConnectionBandwidthQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthQuotas struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthQuotas", string(data)}, " ")
}

type GlobalConnectionBandwidthQuotasType struct {
	value string
}

type GlobalConnectionBandwidthQuotasTypeEnum struct {
	GCB_SIZE  GlobalConnectionBandwidthQuotasType
	GCB_COUNT GlobalConnectionBandwidthQuotasType
}

func GetGlobalConnectionBandwidthQuotasTypeEnum() GlobalConnectionBandwidthQuotasTypeEnum {
	return GlobalConnectionBandwidthQuotasTypeEnum{
		GCB_SIZE: GlobalConnectionBandwidthQuotasType{
			value: "gcb.size",
		},
		GCB_COUNT: GlobalConnectionBandwidthQuotasType{
			value: "gcb.count",
		},
	}
}

func (c GlobalConnectionBandwidthQuotasType) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthQuotasType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthQuotasType) UnmarshalJSON(b []byte) error {
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
