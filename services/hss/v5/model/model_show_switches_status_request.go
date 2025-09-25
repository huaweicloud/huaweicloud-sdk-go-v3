package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSwitchesStatusRequest Request Object
type ShowSwitchesStatusRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 配置类型   - malware_sample_high_detect： 恶意软件高检出模式   - image_pay_per_scan： 镜像扫描按次计费功能开关   - image_popup： 镜像扫描按次计费弹窗开关   - image_free_to_pay_popup：镜像扫描功能免费转付费弹窗开关   - display_unprotected_host：显示未防护的主机
	Code ShowSwitchesStatusRequestCode `json:"code"`
}

func (o ShowSwitchesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSwitchesStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSwitchesStatusRequest", string(data)}, " ")
}

type ShowSwitchesStatusRequestCode struct {
	value string
}

type ShowSwitchesStatusRequestCodeEnum struct {
	MALWARE_SAMPLE_HIGH_DETECT ShowSwitchesStatusRequestCode
	IMAGE_PAY_PER_SCAN         ShowSwitchesStatusRequestCode
	IMAGE_POPUP                ShowSwitchesStatusRequestCode
	IMAGE_FREE_TO_PAY_POPUP    ShowSwitchesStatusRequestCode
}

func GetShowSwitchesStatusRequestCodeEnum() ShowSwitchesStatusRequestCodeEnum {
	return ShowSwitchesStatusRequestCodeEnum{
		MALWARE_SAMPLE_HIGH_DETECT: ShowSwitchesStatusRequestCode{
			value: "malware_sample_high_detect",
		},
		IMAGE_PAY_PER_SCAN: ShowSwitchesStatusRequestCode{
			value: "image_pay_per_scan",
		},
		IMAGE_POPUP: ShowSwitchesStatusRequestCode{
			value: "image_popup",
		},
		IMAGE_FREE_TO_PAY_POPUP: ShowSwitchesStatusRequestCode{
			value: "image_free_to_pay_popup",
		},
	}
}

func (c ShowSwitchesStatusRequestCode) Value() string {
	return c.value
}

func (c ShowSwitchesStatusRequestCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSwitchesStatusRequestCode) UnmarshalJSON(b []byte) error {
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
