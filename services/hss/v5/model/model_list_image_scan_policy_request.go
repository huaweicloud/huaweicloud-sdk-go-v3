package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListImageScanPolicyRequest Request Object
type ListImageScanPolicyRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像类型:   - local：本地镜像   - registry：仓库镜像
	GlobalImageType *string `json:"global_image_type,omitempty"`

	// 策略类型
	Type ListImageScanPolicyRequestType `json:"type"`
}

func (o ListImageScanPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageScanPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListImageScanPolicyRequest", string(data)}, " ")
}

type ListImageScanPolicyRequestType struct {
	value string
}

type ListImageScanPolicyRequestTypeEnum struct {
	CYCLE  ListImageScanPolicyRequestType
	MANUAL ListImageScanPolicyRequestType
}

func GetListImageScanPolicyRequestTypeEnum() ListImageScanPolicyRequestTypeEnum {
	return ListImageScanPolicyRequestTypeEnum{
		CYCLE: ListImageScanPolicyRequestType{
			value: "cycle",
		},
		MANUAL: ListImageScanPolicyRequestType{
			value: "manual",
		},
	}
}

func (c ListImageScanPolicyRequestType) Value() string {
	return c.value
}

func (c ListImageScanPolicyRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageScanPolicyRequestType) UnmarshalJSON(b []byte) error {
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
