package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UploadKieRequest Request Object
type UploadKieRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 微服务引擎ID。
	XEngineId string `json:"x-engine-id"`

	// 覆盖策略，force 强制覆盖、abort 遇到第一个重复时终止导入后续的kv、skip 跳过重复的key
	Override UploadKieRequestOverride `json:"override"`

	// 指定label导入，格式为：{标签key}:{标签value}，如果不填则按body的label导入
	Label *string `json:"label,omitempty"`

	Body *UploadKieRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadKieRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKieRequest struct{}"
	}

	return strings.Join([]string{"UploadKieRequest", string(data)}, " ")
}

type UploadKieRequestOverride struct {
	value string
}

type UploadKieRequestOverrideEnum struct {
	FORCE UploadKieRequestOverride
	ABORT UploadKieRequestOverride
	SKIP  UploadKieRequestOverride
}

func GetUploadKieRequestOverrideEnum() UploadKieRequestOverrideEnum {
	return UploadKieRequestOverrideEnum{
		FORCE: UploadKieRequestOverride{
			value: "force",
		},
		ABORT: UploadKieRequestOverride{
			value: "abort",
		},
		SKIP: UploadKieRequestOverride{
			value: "skip",
		},
	}
}

func (c UploadKieRequestOverride) Value() string {
	return c.value
}

func (c UploadKieRequestOverride) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadKieRequestOverride) UnmarshalJSON(b []byte) error {
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
