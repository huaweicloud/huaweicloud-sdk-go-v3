package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPackageProductsRequest Request Object
type ListPackageProductsRequest struct {

	// **参数解释**： 参数表示用户的语言/所在区域。根据 locale 参数，系统会返回适合该语言/区域的套餐包名称。 **约束限制：** 不涉及 **取值范围**： - zh-cn: 显示中文名称，例如：“Autopilot 通用型 1,000 核时CPU月包” - en-us: 显示英文名称，例如：“Autopilot General Computing 1,000 vCPU-hours CPU monthly package”  **默认取值：** zh-cn
	Locale *ListPackageProductsRequestLocale `json:"locale,omitempty"`
}

func (o ListPackageProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackageProductsRequest struct{}"
	}

	return strings.Join([]string{"ListPackageProductsRequest", string(data)}, " ")
}

type ListPackageProductsRequestLocale struct {
	value string
}

type ListPackageProductsRequestLocaleEnum struct {
	ZH_CN ListPackageProductsRequestLocale
	EN_US ListPackageProductsRequestLocale
}

func GetListPackageProductsRequestLocaleEnum() ListPackageProductsRequestLocaleEnum {
	return ListPackageProductsRequestLocaleEnum{
		ZH_CN: ListPackageProductsRequestLocale{
			value: "zh-cn",
		},
		EN_US: ListPackageProductsRequestLocale{
			value: "en-us",
		},
	}
}

func (c ListPackageProductsRequestLocale) Value() string {
	return c.value
}

func (c ListPackageProductsRequestLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPackageProductsRequestLocale) UnmarshalJSON(b []byte) error {
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
