package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGlobalConnectionBandwidthLineLevelsRequest Request Object
type ListGlobalConnectionBandwidthLineLevelsRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向后翻页。 翻页过程中，查询条件不能修改，包括过滤条件、排序条件、limit。
	Marker *string `json:"marker,omitempty"`

	// 根据ID查询，可查询多个ID。
	Id *[]string `json:"id,omitempty"`

	// 线路规格本端接入点编码信息。
	LocalArea *string `json:"local_area,omitempty"`

	// 线路规格远端接入点编码信息。
	RemoteArea *string `json:"remote_area,omitempty"`

	// 带宽等级信息： - Pt: 铂金 - Ag: 银
	Levels *[]ListGlobalConnectionBandwidthLineLevelsRequestLevels `json:"levels,omitempty"`
}

func (o ListGlobalConnectionBandwidthLineLevelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthLineLevelsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthLineLevelsRequest", string(data)}, " ")
}

type ListGlobalConnectionBandwidthLineLevelsRequestLevels struct {
	value string
}

type ListGlobalConnectionBandwidthLineLevelsRequestLevelsEnum struct {
	PT ListGlobalConnectionBandwidthLineLevelsRequestLevels
	AG ListGlobalConnectionBandwidthLineLevelsRequestLevels
}

func GetListGlobalConnectionBandwidthLineLevelsRequestLevelsEnum() ListGlobalConnectionBandwidthLineLevelsRequestLevelsEnum {
	return ListGlobalConnectionBandwidthLineLevelsRequestLevelsEnum{
		PT: ListGlobalConnectionBandwidthLineLevelsRequestLevels{
			value: "Pt",
		},
		AG: ListGlobalConnectionBandwidthLineLevelsRequestLevels{
			value: "Ag",
		},
	}
}

func (c ListGlobalConnectionBandwidthLineLevelsRequestLevels) Value() string {
	return c.value
}

func (c ListGlobalConnectionBandwidthLineLevelsRequestLevels) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGlobalConnectionBandwidthLineLevelsRequestLevels) UnmarshalJSON(b []byte) error {
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
