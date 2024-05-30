package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EngineCreateReq 创建微服务引擎请求结构体
type EngineCreateReq struct {

	// 微服务引擎的名称，名称为字母开头，字母、数字、-组成，且不能以-结尾，3-24个字符。
	Name string `json:"name"`

	// 微服务引擎描述，长度0~255。
	Description *string `json:"description,omitempty"`

	// 微服务引擎计费方式，1表示按需
	Payment EngineCreateReqPayment `json:"payment"`

	// 微服务引擎的规格
	Flavor EngineCreateReqFlavor `json:"flavor"`

	// 当前局点可用区列表，创建ServiceComb引擎专享版需要填写。
	AzList *[]string `json:"azList,omitempty"`

	// ServiceComb引擎专享版与注册配置中心认证方式，RBAC为安全认证，NONE为无认证。
	AuthType EngineCreateReqAuthType `json:"authType"`

	// vpc名称
	Vpc string `json:"vpc"`

	// vpc标识
	VpcId *string `json:"vpcId,omitempty"`

	// 微服务引擎子网ID
	NetworkId string `json:"networkId"`

	// 微服务引擎子网划分
	SubnetCidr string `json:"subnetCidr"`

	// ServiceComb引擎专享版公网地址ID，当前为null
	PublicIpId *string `json:"publicIpId,omitempty"`

	AuthCred *EngineRbacPwd `json:"auth_cred,omitempty"`

	// 微服务引擎部署类型
	SpecType EngineCreateReqSpecType `json:"specType"`

	// 引擎附加参数
	Inputs map[string]string `json:"inputs,omitempty"`

	EnginestateInfo *EngineCreateReqEnginestateInfo `json:"enginestateInfo,omitempty"`

	// 创建阶段类型
	PeriodType *int32 `json:"periodType,omitempty"`

	FlavorType *EngineCreateReqFlavorType `json:"flavorType,omitempty"`

	EnterpriseProject *EngineCreateReqEnterpriseProject `json:"enterpriseProject,omitempty"`

	// 网关vpc划分
	VpcCidr *string `json:"vpcCidr,omitempty"`

	ResourceParams *EngineCreateReqResourceParams `json:"resourceParams,omitempty"`

	// 产品ID
	ProductId *string `json:"productId,omitempty"`

	// 容量产品ID
	CapacityProductId *string `json:"capacityProductId,omitempty"`

	// 微服务引擎是否免费
	IsFree *bool `json:"isFree,omitempty"`

	// 微服务引擎使用的子网名称
	SubnetName *string `json:"subnetName,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	MaintenanceConfig *EngineCreateReqMaintenanceConfig `json:"maintenanceConfig,omitempty"`

	// 微服务引擎使用的elb的id
	Elbid *string `json:"elbid,omitempty"`
}

func (o EngineCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReq struct{}"
	}

	return strings.Join([]string{"EngineCreateReq", string(data)}, " ")
}

type EngineCreateReqPayment struct {
	value string
}

type EngineCreateReqPaymentEnum struct {
	E_1 EngineCreateReqPayment
}

func GetEngineCreateReqPaymentEnum() EngineCreateReqPaymentEnum {
	return EngineCreateReqPaymentEnum{
		E_1: EngineCreateReqPayment{
			value: "1",
		},
	}
}

func (c EngineCreateReqPayment) Value() string {
	return c.value
}

func (c EngineCreateReqPayment) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqPayment) UnmarshalJSON(b []byte) error {
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

type EngineCreateReqFlavor struct {
	value string
}

type EngineCreateReqFlavorEnum struct {
	CSE_S1_SMALL2                 EngineCreateReqFlavor
	CSE_S1_MEDIUM2                EngineCreateReqFlavor
	CSE_S1_LARGE2                 EngineCreateReqFlavor
	CSE_S1_XLARGE2                EngineCreateReqFlavor
	CSE_NACOS2_C1_LARGE_10        EngineCreateReqFlavor
	CSE_NACOS2_C1_XLARGE_20       EngineCreateReqFlavor
	CSE_NACOS2_C1_XLARGE_50       EngineCreateReqFlavor
	CSE_NACOS2_C1_XLARGE_60       EngineCreateReqFlavor
	CSE_NACOS2_C1_2XLARGE_100     EngineCreateReqFlavor
	CSE_MICROGATEWAY_PRO_SMALL_1  EngineCreateReqFlavor
	CSE_MICROGATEWAY_PRO_MEDIUM_1 EngineCreateReqFlavor
	CSE_MICROGATEWAY_PRO_LARGE_1  EngineCreateReqFlavor
}

func GetEngineCreateReqFlavorEnum() EngineCreateReqFlavorEnum {
	return EngineCreateReqFlavorEnum{
		CSE_S1_SMALL2: EngineCreateReqFlavor{
			value: "cse.s1.small2",
		},
		CSE_S1_MEDIUM2: EngineCreateReqFlavor{
			value: "cse.s1.medium2",
		},
		CSE_S1_LARGE2: EngineCreateReqFlavor{
			value: "cse.s1.large2",
		},
		CSE_S1_XLARGE2: EngineCreateReqFlavor{
			value: "cse.s1.xlarge2",
		},
		CSE_NACOS2_C1_LARGE_10: EngineCreateReqFlavor{
			value: "cse.nacos2.c1.large.10",
		},
		CSE_NACOS2_C1_XLARGE_20: EngineCreateReqFlavor{
			value: "cse.nacos2.c1.xlarge.20",
		},
		CSE_NACOS2_C1_XLARGE_50: EngineCreateReqFlavor{
			value: "cse.nacos2.c1.xlarge.50",
		},
		CSE_NACOS2_C1_XLARGE_60: EngineCreateReqFlavor{
			value: "cse.nacos2.c1.xlarge.60",
		},
		CSE_NACOS2_C1_2XLARGE_100: EngineCreateReqFlavor{
			value: "cse.nacos2.c1.2xlarge.100",
		},
		CSE_MICROGATEWAY_PRO_SMALL_1: EngineCreateReqFlavor{
			value: "cse.microgateway.pro.small.1",
		},
		CSE_MICROGATEWAY_PRO_MEDIUM_1: EngineCreateReqFlavor{
			value: "cse.microgateway.pro.medium.1",
		},
		CSE_MICROGATEWAY_PRO_LARGE_1: EngineCreateReqFlavor{
			value: "cse.microgateway.pro.large.1",
		},
	}
}

func (c EngineCreateReqFlavor) Value() string {
	return c.value
}

func (c EngineCreateReqFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqFlavor) UnmarshalJSON(b []byte) error {
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

type EngineCreateReqAuthType struct {
	value string
}

type EngineCreateReqAuthTypeEnum struct {
	RBAC EngineCreateReqAuthType
	NONE EngineCreateReqAuthType
}

func GetEngineCreateReqAuthTypeEnum() EngineCreateReqAuthTypeEnum {
	return EngineCreateReqAuthTypeEnum{
		RBAC: EngineCreateReqAuthType{
			value: "RBAC",
		},
		NONE: EngineCreateReqAuthType{
			value: "NONE",
		},
	}
}

func (c EngineCreateReqAuthType) Value() string {
	return c.value
}

func (c EngineCreateReqAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqAuthType) UnmarshalJSON(b []byte) error {
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

type EngineCreateReqSpecType struct {
	value string
}

type EngineCreateReqSpecTypeEnum struct {
	CSE2          EngineCreateReqSpecType
	NACOS2        EngineCreateReqSpecType
	MICRO_GATEWAY EngineCreateReqSpecType
}

func GetEngineCreateReqSpecTypeEnum() EngineCreateReqSpecTypeEnum {
	return EngineCreateReqSpecTypeEnum{
		CSE2: EngineCreateReqSpecType{
			value: "CSE2",
		},
		NACOS2: EngineCreateReqSpecType{
			value: "Nacos2",
		},
		MICRO_GATEWAY: EngineCreateReqSpecType{
			value: "MicroGateway",
		},
	}
}

func (c EngineCreateReqSpecType) Value() string {
	return c.value
}

func (c EngineCreateReqSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqSpecType) UnmarshalJSON(b []byte) error {
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
