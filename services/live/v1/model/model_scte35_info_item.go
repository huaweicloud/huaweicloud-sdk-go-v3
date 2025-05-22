package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Scte35InfoItem SCTE35信号信息的item结构。
type Scte35InfoItem struct {

	// 信号类型，splice_insert/time_signal。
	Type Scte35InfoItemType `json:"type"`

	// 广告信号的Event ID，Time Signal打印数组第一个。
	EventId int32 `json:"event_id"`

	// 广告信号的执行时间，unix time，单位：秒。
	StartDate int64 `json:"start_date"`

	// 广告信号时长，-1表示没有携带,单位：秒。
	Duration int32 `json:"duration"`

	// // Splice Insert填空\"-\"； // Time Signal，支持0x30，0x31，0x32，0x33，0x34，0x35，0x36，0x37 // 0x30: ProviderAdvertisementStart // 0x31: ProviderAdvertisementEnd // 0x32: DistributorAdvertisementStart // 0x33: DistributorAdvertisementEnd // 0x34: ProviderPlacementOpportunityStart // 0x35: ProviderPlacementOpportunityEnd // 0x36: DistributorPlacementOpportunityStart // 0x37: DistributorPlacementOpportunityEnd
	SegmentationType Scte35InfoItemSegmentationType `json:"segmentation_type"`

	// 广告信号原始数据的base64值。
	Base64Data string `json:"base64_data"`

	// 广告信号全量信息。
	RawSplice string `json:"raw_splice"`
}

func (o Scte35InfoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scte35InfoItem struct{}"
	}

	return strings.Join([]string{"Scte35InfoItem", string(data)}, " ")
}

type Scte35InfoItemType struct {
	value string
}

type Scte35InfoItemTypeEnum struct {
	SPLICE_INSERT Scte35InfoItemType
	TIME_SIGNAL   Scte35InfoItemType
}

func GetScte35InfoItemTypeEnum() Scte35InfoItemTypeEnum {
	return Scte35InfoItemTypeEnum{
		SPLICE_INSERT: Scte35InfoItemType{
			value: "splice_insert",
		},
		TIME_SIGNAL: Scte35InfoItemType{
			value: "time_signal",
		},
	}
}

func (c Scte35InfoItemType) Value() string {
	return c.value
}

func (c Scte35InfoItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Scte35InfoItemType) UnmarshalJSON(b []byte) error {
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

type Scte35InfoItemSegmentationType struct {
	value string
}

type Scte35InfoItemSegmentationTypeEnum struct {
	MINUS                                   Scte35InfoItemSegmentationType
	PROVIDER_ADVERTISEMENT_START            Scte35InfoItemSegmentationType
	PROVIDER_ADVERTISEMENT_END              Scte35InfoItemSegmentationType
	DISTRIBUTOR_ADVERTISEMENT_START         Scte35InfoItemSegmentationType
	DISTRIBUTOR_ADVERTISEMENT_END           Scte35InfoItemSegmentationType
	PROVIDER_PLACEMENT_OPPORTUNITY_START    Scte35InfoItemSegmentationType
	PROVIDER_PLACEMENT_OPPORTUNITY_END      Scte35InfoItemSegmentationType
	DISTRIBUTOR_PLACEMENT_OPPORTUNITY_START Scte35InfoItemSegmentationType
	DISTRIBUTOR_PLACEMENT_OPPORTUNITY_END   Scte35InfoItemSegmentationType
}

func GetScte35InfoItemSegmentationTypeEnum() Scte35InfoItemSegmentationTypeEnum {
	return Scte35InfoItemSegmentationTypeEnum{
		MINUS: Scte35InfoItemSegmentationType{
			value: "-",
		},
		PROVIDER_ADVERTISEMENT_START: Scte35InfoItemSegmentationType{
			value: "ProviderAdvertisementStart",
		},
		PROVIDER_ADVERTISEMENT_END: Scte35InfoItemSegmentationType{
			value: "ProviderAdvertisementEnd",
		},
		DISTRIBUTOR_ADVERTISEMENT_START: Scte35InfoItemSegmentationType{
			value: "DistributorAdvertisementStart",
		},
		DISTRIBUTOR_ADVERTISEMENT_END: Scte35InfoItemSegmentationType{
			value: "DistributorAdvertisementEnd",
		},
		PROVIDER_PLACEMENT_OPPORTUNITY_START: Scte35InfoItemSegmentationType{
			value: "ProviderPlacementOpportunityStart",
		},
		PROVIDER_PLACEMENT_OPPORTUNITY_END: Scte35InfoItemSegmentationType{
			value: "ProviderPlacementOpportunityEnd",
		},
		DISTRIBUTOR_PLACEMENT_OPPORTUNITY_START: Scte35InfoItemSegmentationType{
			value: "DistributorPlacementOpportunityStart",
		},
		DISTRIBUTOR_PLACEMENT_OPPORTUNITY_END: Scte35InfoItemSegmentationType{
			value: "DistributorPlacementOpportunityEnd",
		},
	}
}

func (c Scte35InfoItemSegmentationType) Value() string {
	return c.value
}

func (c Scte35InfoItemSegmentationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Scte35InfoItemSegmentationType) UnmarshalJSON(b []byte) error {
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
