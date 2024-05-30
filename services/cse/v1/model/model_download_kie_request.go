package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadKieRequest Request Object
type DownloadKieRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 微服务引擎ID。
	XEngineId string `json:"x-engine-id"`

	// 按label过滤项导出，格式为：{标签key}:{标签value}
	Label *string `json:"label,omitempty"`

	// 对label过滤项的匹配选项，如果值为exact：表示严格匹配，包括label个数和内容相等；不填表示包含匹配
	Match *DownloadKieRequestMatch `json:"match,omitempty"`

	Body *DownloadKieReqBody `json:"body,omitempty"`
}

func (o DownloadKieRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieRequest struct{}"
	}

	return strings.Join([]string{"DownloadKieRequest", string(data)}, " ")
}

type DownloadKieRequestMatch struct {
	value string
}

type DownloadKieRequestMatchEnum struct {
	EXACT DownloadKieRequestMatch
}

func GetDownloadKieRequestMatchEnum() DownloadKieRequestMatchEnum {
	return DownloadKieRequestMatchEnum{
		EXACT: DownloadKieRequestMatch{
			value: "exact",
		},
	}
}

func (c DownloadKieRequestMatch) Value() string {
	return c.value
}

func (c DownloadKieRequestMatch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadKieRequestMatch) UnmarshalJSON(b []byte) error {
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
