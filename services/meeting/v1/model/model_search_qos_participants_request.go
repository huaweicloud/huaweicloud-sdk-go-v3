package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SearchQosParticipantsRequest struct {

	// 会议UUID。最大不超过64个字节。
	ConfUUID string `json:"confUUID"`

	// 会议类别。 * online：在线会议，在召开的会议。 * history：历史会议，已召开的会议。
	ConfType SearchQosParticipantsRequestConfType `json:"confType"`

	// 查询偏移量。 * 取值：大于等于0，默认值为0。 * 小于最小值0时，系统设置为0。 * 大于等于最大条目数量，则返回最后一页的数据。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20。 * 小于最小值1时，系统设置为1。 * 大于最大值500时，系统设置为500。
	Limit *int32 `json:"limit,omitempty"`

	// 根据与会人名称作为关键词，模糊查询与会者列表
	SearchKey *string `json:"searchKey,omitempty"`
}

func (o SearchQosParticipantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosParticipantsRequest struct{}"
	}

	return strings.Join([]string{"SearchQosParticipantsRequest", string(data)}, " ")
}

type SearchQosParticipantsRequestConfType struct {
	value string
}

type SearchQosParticipantsRequestConfTypeEnum struct {
	ONLINE  SearchQosParticipantsRequestConfType
	HISTORY SearchQosParticipantsRequestConfType
}

func GetSearchQosParticipantsRequestConfTypeEnum() SearchQosParticipantsRequestConfTypeEnum {
	return SearchQosParticipantsRequestConfTypeEnum{
		ONLINE: SearchQosParticipantsRequestConfType{
			value: "online",
		},
		HISTORY: SearchQosParticipantsRequestConfType{
			value: "history",
		},
	}
}

func (c SearchQosParticipantsRequestConfType) Value() string {
	return c.value
}

func (c SearchQosParticipantsRequestConfType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchQosParticipantsRequestConfType) UnmarshalJSON(b []byte) error {
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
