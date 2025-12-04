package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CustomPage 自定义告警页面
type CustomPage struct {

	// 返回状态码
	StatusCode string `json:"status_code"`

	// **参数解释：** “自定义”告警页面内容类型 **约束限制：** 不涉及 **取值范围：**  - text/html  - text/xml  - application/json  **默认取值：** 不涉及
	ContentType CustomPageContentType `json:"content_type"`

	// 根据选择的“页面类型”配置对应的页面内容，具体示例可以参考“Web应用防火墙 WAF”用户手册
	Content string `json:"content"`
}

func (o CustomPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPage struct{}"
	}

	return strings.Join([]string{"CustomPage", string(data)}, " ")
}

type CustomPageContentType struct {
	value string
}

type CustomPageContentTypeEnum struct {
	TEXT_HTML        CustomPageContentType
	TEXT_XML         CustomPageContentType
	APPLICATION_JSON CustomPageContentType
}

func GetCustomPageContentTypeEnum() CustomPageContentTypeEnum {
	return CustomPageContentTypeEnum{
		TEXT_HTML: CustomPageContentType{
			value: "text/html",
		},
		TEXT_XML: CustomPageContentType{
			value: "text/xml",
		},
		APPLICATION_JSON: CustomPageContentType{
			value: "application/json",
		},
	}
}

func (c CustomPageContentType) Value() string {
	return c.value
}

func (c CustomPageContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomPageContentType) UnmarshalJSON(b []byte) error {
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
