package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowJobsRequest struct {

	// 接收数据类型，支持两种接收数据类型：“管道数据”、“资产数据”。管道数据：“实时分析”使用来自“数据管道”的数据进行分析，并可将数据输出到其他云服务。资产数据：“实时分析”使用来自“资产模型”的数据进行分析，并将分析后的结果返回给“资产模型”，丰富资产模型。
	JobInputType *ShowJobsRequestJobInputType `json:"job_input_type,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 立即同步作业状态，默认是false
	SyncStatus *bool `json:"sync_status,omitempty"`
}

func (o ShowJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobsRequest struct{}"
	}

	return strings.Join([]string{"ShowJobsRequest", string(data)}, " ")
}

type ShowJobsRequestJobInputType struct {
	value string
}

type ShowJobsRequestJobInputTypeEnum struct {
	ASSET_DATA ShowJobsRequestJobInputType
	OLD_DATA   ShowJobsRequestJobInputType
}

func GetShowJobsRequestJobInputTypeEnum() ShowJobsRequestJobInputTypeEnum {
	return ShowJobsRequestJobInputTypeEnum{
		ASSET_DATA: ShowJobsRequestJobInputType{
			value: "ASSET_DATA",
		},
		OLD_DATA: ShowJobsRequestJobInputType{
			value: "OLD_DATA",
		},
	}
}

func (c ShowJobsRequestJobInputType) Value() string {
	return c.value
}

func (c ShowJobsRequestJobInputType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobsRequestJobInputType) UnmarshalJSON(b []byte) error {
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
