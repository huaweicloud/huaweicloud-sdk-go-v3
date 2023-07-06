package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BandwidthInfoResp 弹性公网IP绑定的带宽信息
type BandwidthInfoResp struct {

	// - 功能说明：带宽名称
	BandwidthName *string `json:"bandwidth_name,omitempty"`

	// - 功能说明：带宽大小
	BandwidthNumber *int32 `json:"bandwidth_number,omitempty"`

	// - 功能说明：带宽类型
	BandwidthType *BandwidthInfoRespBandwidthType `json:"bandwidth_type,omitempty"`

	// - 功能说明：带宽id
	BandwidthId *string `json:"bandwidth_id,omitempty"`
}

func (o BandwidthInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthInfoResp struct{}"
	}

	return strings.Join([]string{"BandwidthInfoResp", string(data)}, " ")
}

type BandwidthInfoRespBandwidthType struct {
	value string
}

type BandwidthInfoRespBandwidthTypeEnum struct {
	PER   BandwidthInfoRespBandwidthType
	WHOLE BandwidthInfoRespBandwidthType
}

func GetBandwidthInfoRespBandwidthTypeEnum() BandwidthInfoRespBandwidthTypeEnum {
	return BandwidthInfoRespBandwidthTypeEnum{
		PER: BandwidthInfoRespBandwidthType{
			value: "PER",
		},
		WHOLE: BandwidthInfoRespBandwidthType{
			value: "WHOLE",
		},
	}
}

func (c BandwidthInfoRespBandwidthType) Value() string {
	return c.value
}

func (c BandwidthInfoRespBandwidthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthInfoRespBandwidthType) UnmarshalJSON(b []byte) error {
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
