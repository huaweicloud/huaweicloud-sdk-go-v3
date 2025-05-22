package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodePoolUpdateExtendParam **参数解释**： 节点池更新时支持的扩展参数。 **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type NodePoolUpdateExtendParam struct {

	// **参数解释**： 委托的名称。 委托是由租户管理员在统一身份认证服务（Identity and AccessManagement，IAM）上创建的，可以为CCE节点提供访问云服务器的临时凭证。 **约束限制**： 作为响应参数仅在创建节点传入时返回该字段。 **取值范围：** 不涉及 **默认取值：** 不涉及
	AgencyName *string `json:"agency_name,omitempty"`

	// **参数解释**： 安装前执行脚本。 输入的值需要经过Base64编码，方法如下：   ```   echo -n \"待编码内容\" | base64   ```   **约束限制**：  安装前/后执行脚本统一计算字符，转码后的字符总数不能超过10240。  **取值范围：**  不涉及  **默认取值：**  不涉及
	AlphaCcePreInstall *string `json:"alpha.cce/preInstall,omitempty"`

	// **参数解释**： 安装后执行脚本。 输入的值需要经过Base64编码，方法如下：   ```   echo -n \"待编码内容\" | base64   ```  **约束限制**： 安装前/后执行脚本统一计算字符，转码后的字符总数不能超过10240。 **取值范围：** 不涉及 **默认取值：** 不涉及
	AlphaCcePostInstall *string `json:"alpha.cce/postInstall,omitempty"`

	// **参数解释**： 用户愿意为竞价实例每小时支付的最高价格。 **约束限制**： - 仅billingMode=0且marketType=spot时，该参数设置后生效。 - 当billingMode=0且marketType=spot时，如果不传递spotPrice，默认使用按需购买的价格作为竞价。 - spotPrice需要小于等于按需价格并大于等于云服务器市场价格。  **取值范围：** 不涉及 **默认取值：** 不涉及
	SpotPrice *string `json:"spotPrice,omitempty"`

	// **参数解释**： 指定节点安全加固类型，当前仅支持HCE2.0镜像等保2.0三级安全加固。 等保加固会对身份鉴别、访问控制、安全审计、入侵防范、恶意代码防范进行检查并加固。[详情请参见[Huawei Cloud EulerOS 2.0等保2.0三级版镜像概述](https://support.huaweicloud.com/productdesc-hce/hce_sec_0001.html)。](tag:hws) 若未指定此参数，则尝试用原有的值补全。如：原先HCE2.0镜像已配置安全加固，更新节点池时未指定此参数，则仍旧保持安全加固配置，若要取消，需显式指定参数值为\"null\"。 **约束限制**： 不涉及 **取值范围**： - 空值：表示不开启等保加固 - cybersecurity：表示开启等保加固  **默认取值**： 不涉及
	SecurityReinforcementType *NodePoolUpdateExtendParamSecurityReinforcementType `json:"securityReinforcementType,omitempty"`
}

func (o NodePoolUpdateExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolUpdateExtendParam struct{}"
	}

	return strings.Join([]string{"NodePoolUpdateExtendParam", string(data)}, " ")
}

type NodePoolUpdateExtendParamSecurityReinforcementType struct {
	value string
}

type NodePoolUpdateExtendParamSecurityReinforcementTypeEnum struct {
	NULL          NodePoolUpdateExtendParamSecurityReinforcementType
	CYBERSECURITY NodePoolUpdateExtendParamSecurityReinforcementType
}

func GetNodePoolUpdateExtendParamSecurityReinforcementTypeEnum() NodePoolUpdateExtendParamSecurityReinforcementTypeEnum {
	return NodePoolUpdateExtendParamSecurityReinforcementTypeEnum{
		NULL: NodePoolUpdateExtendParamSecurityReinforcementType{
			value: "null",
		},
		CYBERSECURITY: NodePoolUpdateExtendParamSecurityReinforcementType{
			value: "cybersecurity",
		},
	}
}

func (c NodePoolUpdateExtendParamSecurityReinforcementType) Value() string {
	return c.value
}

func (c NodePoolUpdateExtendParamSecurityReinforcementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodePoolUpdateExtendParamSecurityReinforcementType) UnmarshalJSON(b []byte) error {
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
