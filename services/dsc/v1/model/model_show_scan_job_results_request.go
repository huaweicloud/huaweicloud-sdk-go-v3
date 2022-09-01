package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowScanJobResultsRequest struct {

	// 任务ID
	JobId string `json:"job_id" xml:"job_id"`

	// 页码
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 资产类型
	Type *ShowScanJobResultsRequestType `json:"type,omitempty" xml:"type"`

	// 预留，待启用
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 预留，待启用
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o ShowScanJobResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobResultsRequest struct{}"
	}

	return strings.Join([]string{"ShowScanJobResultsRequest", string(data)}, " ")
}

type ShowScanJobResultsRequestType struct {
	value string
}

type ShowScanJobResultsRequestTypeEnum struct {
	DATABASE ShowScanJobResultsRequestType
	OBS      ShowScanJobResultsRequestType
	BIGDATA  ShowScanJobResultsRequestType
}

func GetShowScanJobResultsRequestTypeEnum() ShowScanJobResultsRequestTypeEnum {
	return ShowScanJobResultsRequestTypeEnum{
		DATABASE: ShowScanJobResultsRequestType{
			value: "DATABASE",
		},
		OBS: ShowScanJobResultsRequestType{
			value: "OBS",
		},
		BIGDATA: ShowScanJobResultsRequestType{
			value: "BIGDATA",
		},
	}
}

func (c ShowScanJobResultsRequestType) Value() string {
	return c.value
}

func (c ShowScanJobResultsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowScanJobResultsRequestType) UnmarshalJSON(b []byte) error {
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
