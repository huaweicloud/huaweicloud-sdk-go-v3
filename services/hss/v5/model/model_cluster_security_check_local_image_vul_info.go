package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClusterSecurityCheckLocalImageVulInfo 本地镜像漏洞信息
type ClusterSecurityCheckLocalImageVulInfo struct {

	// **参数解释**： 本地镜像名称 **取值范围**： 不涉及
	LocalImageName *string `json:"local_image_name,omitempty"`

	// **参数解释**： 本地镜像版本 **取值范围**： 不涉及
	LocalImageVersion *string `json:"local_image_version,omitempty"`

	// **参数解释**： 本地镜像大小 **取值范围**： 不涉及
	LocalImageSize *int64 `json:"local_image_size,omitempty"`

	// **参数解释**： 漏洞名称 **取值范围**： 不涉及
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 软件名称 **取值范围**： 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**： 软件版本 **取值范围**： 不涉及
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**： 漏洞危险程度 **取值范围**： - High：高危漏洞 - Medium：中危漏洞 - Low：低危漏洞
	SeverityLevel *ClusterSecurityCheckLocalImageVulInfoSeverityLevel `json:"severity_level,omitempty"`

	// **参数解释**： 漏洞解决方案 **取值范围**： 不涉及
	Solution *string `json:"solution,omitempty"`

	// **参数解释**： 漏洞描述 **取值范围**： 不涉及
	VulDescription *string `json:"vul_description,omitempty"`
}

func (o ClusterSecurityCheckLocalImageVulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckLocalImageVulInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckLocalImageVulInfo", string(data)}, " ")
}

type ClusterSecurityCheckLocalImageVulInfoSeverityLevel struct {
	value string
}

type ClusterSecurityCheckLocalImageVulInfoSeverityLevelEnum struct {
	HIGH   ClusterSecurityCheckLocalImageVulInfoSeverityLevel
	MEDIUM ClusterSecurityCheckLocalImageVulInfoSeverityLevel
	LOW    ClusterSecurityCheckLocalImageVulInfoSeverityLevel
}

func GetClusterSecurityCheckLocalImageVulInfoSeverityLevelEnum() ClusterSecurityCheckLocalImageVulInfoSeverityLevelEnum {
	return ClusterSecurityCheckLocalImageVulInfoSeverityLevelEnum{
		HIGH: ClusterSecurityCheckLocalImageVulInfoSeverityLevel{
			value: "High",
		},
		MEDIUM: ClusterSecurityCheckLocalImageVulInfoSeverityLevel{
			value: "Medium",
		},
		LOW: ClusterSecurityCheckLocalImageVulInfoSeverityLevel{
			value: "Low",
		},
	}
}

func (c ClusterSecurityCheckLocalImageVulInfoSeverityLevel) Value() string {
	return c.value
}

func (c ClusterSecurityCheckLocalImageVulInfoSeverityLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSecurityCheckLocalImageVulInfoSeverityLevel) UnmarshalJSON(b []byte) error {
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
