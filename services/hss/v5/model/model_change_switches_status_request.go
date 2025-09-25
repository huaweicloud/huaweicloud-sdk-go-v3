package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeSwitchesStatusRequest Request Object
type ChangeSwitchesStatusRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 配置类型   - malware_sample_high_detect： 恶意软件高检出模式   - image_pay_per_scan： 镜像扫描按次计费功能开关   - image_popup： 镜像扫描按次计费弹窗开关   - image_free_to_pay_popup：镜像扫描功能免费转付费弹窗开关   - display_unprotected_host：显示未防护的主机
	Code ChangeSwitchesStatusRequestCode `json:"code"`

	Body *ChangeSwitchesStatusRequestBody `json:"body,omitempty"`
}

func (o ChangeSwitchesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSwitchesStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeSwitchesStatusRequest", string(data)}, " ")
}

type ChangeSwitchesStatusRequestCode struct {
	value string
}

type ChangeSwitchesStatusRequestCodeEnum struct {
	MALWARE_SAMPLE_HIGH_DETECT ChangeSwitchesStatusRequestCode
	IMAGE_PAY_PER_SCAN         ChangeSwitchesStatusRequestCode
	IMAGE_POPUP                ChangeSwitchesStatusRequestCode
	IMAGE_FREE_TO_PAY_POPUP    ChangeSwitchesStatusRequestCode
}

func GetChangeSwitchesStatusRequestCodeEnum() ChangeSwitchesStatusRequestCodeEnum {
	return ChangeSwitchesStatusRequestCodeEnum{
		MALWARE_SAMPLE_HIGH_DETECT: ChangeSwitchesStatusRequestCode{
			value: "malware_sample_high_detect",
		},
		IMAGE_PAY_PER_SCAN: ChangeSwitchesStatusRequestCode{
			value: "image_pay_per_scan",
		},
		IMAGE_POPUP: ChangeSwitchesStatusRequestCode{
			value: "image_popup",
		},
		IMAGE_FREE_TO_PAY_POPUP: ChangeSwitchesStatusRequestCode{
			value: "image_free_to_pay_popup",
		},
	}
}

func (c ChangeSwitchesStatusRequestCode) Value() string {
	return c.value
}

func (c ChangeSwitchesStatusRequestCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeSwitchesStatusRequestCode) UnmarshalJSON(b []byte) error {
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
