package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadCompareResultFileRequest Request Object
type DownloadCompareResultFileRequest struct {

	// 请求语言类型。
	XLanguage *DownloadCompareResultFileRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务类型： - contents： 内容对比。 - lines：行数对比。 - random：抽样对比。 - objects_comparison：对象对比。
	CompareType *string `json:"compare_type,omitempty"`

	// 对比任务的ID，内容对比、抽样对比、行数对比场景必填。
	CompareJobId *string `json:"compare_job_id,omitempty"`

	// 区域ID，例如：cn-north-4。
	Region *string `json:"Region,omitempty"`

	Body *ExportCompareResultReq `json:"body,omitempty"`
}

func (o DownloadCompareResultFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCompareResultFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadCompareResultFileRequest", string(data)}, " ")
}

type DownloadCompareResultFileRequestXLanguage struct {
	value string
}

type DownloadCompareResultFileRequestXLanguageEnum struct {
	EN_US DownloadCompareResultFileRequestXLanguage
	ZH_CN DownloadCompareResultFileRequestXLanguage
}

func GetDownloadCompareResultFileRequestXLanguageEnum() DownloadCompareResultFileRequestXLanguageEnum {
	return DownloadCompareResultFileRequestXLanguageEnum{
		EN_US: DownloadCompareResultFileRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DownloadCompareResultFileRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DownloadCompareResultFileRequestXLanguage) Value() string {
	return c.value
}

func (c DownloadCompareResultFileRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadCompareResultFileRequestXLanguage) UnmarshalJSON(b []byte) error {
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
