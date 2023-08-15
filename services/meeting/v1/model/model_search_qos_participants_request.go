package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchQosParticipantsRequest Request Object
type SearchQosParticipantsRequest struct {

	// 会议UUID。
	ConfUUID string `json:"confUUID"`

	// 会议类别。 * online：在线会议，正在召开的会议 * history：历史会议，已结束的会议
	ConfType SearchQosParticipantsRequestConfType `json:"confType"`

	// 查询偏移量。 * 取值：大于等于0，默认值为0。 * 大于等于最大条目数量，则返回最后一页的数据。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20。
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件。与会者名称可作为搜索内容。长度限制为1-512个字符。
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
