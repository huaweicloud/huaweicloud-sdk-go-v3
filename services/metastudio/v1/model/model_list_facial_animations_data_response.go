package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFacialAnimationsDataResponse Response Object
type ListFacialAnimationsDataResponse struct {

	// csv文件下载地址
	CsvFileDownloadUrl *string `json:"csv_file_download_url,omitempty"`

	// 任务的状态。 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败
	State *ListFacialAnimationsDataResponseState `json:"state,omitempty"`

	// 失败任务描述
	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFacialAnimationsDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFacialAnimationsDataResponse struct{}"
	}

	return strings.Join([]string{"ListFacialAnimationsDataResponse", string(data)}, " ")
}

type ListFacialAnimationsDataResponseState struct {
	value string
}

type ListFacialAnimationsDataResponseStateEnum struct {
	PROCESSING ListFacialAnimationsDataResponseState
	SUCCEED    ListFacialAnimationsDataResponseState
	FAILED     ListFacialAnimationsDataResponseState
}

func GetListFacialAnimationsDataResponseStateEnum() ListFacialAnimationsDataResponseStateEnum {
	return ListFacialAnimationsDataResponseStateEnum{
		PROCESSING: ListFacialAnimationsDataResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ListFacialAnimationsDataResponseState{
			value: "SUCCEED",
		},
		FAILED: ListFacialAnimationsDataResponseState{
			value: "FAILED",
		},
	}
}

func (c ListFacialAnimationsDataResponseState) Value() string {
	return c.value
}

func (c ListFacialAnimationsDataResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFacialAnimationsDataResponseState) UnmarshalJSON(b []byte) error {
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
