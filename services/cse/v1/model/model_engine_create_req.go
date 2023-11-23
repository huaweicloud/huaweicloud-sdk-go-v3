package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EngineCreateReq 创建微服务引擎专享版请求结构体
type EngineCreateReq struct {

	// 微服务引擎专享版的名称，名称为字母开头，字母、数字、-组成，且不能以-结尾，3-24个字符。
	Name string `json:"name"`

	// 微服务引擎专享版描述，长度0~255。
	Description *string `json:"description,omitempty"`

	// 微服务引擎专享版计费方式，1表示按需
	Payment EngineCreateReqPayment `json:"payment"`

	// 微服务引擎专享版的规格
	Flavor EngineCreateReqFlavor `json:"flavor"`

	// 当前局点可用区列表。
	AzList []string `json:"azList"`

	// 微服务引擎专享版认证方式，RBAC为安全认证，NONE为无认证。
	AuthType EngineCreateReqAuthType `json:"authType"`

	// vpc名称
	Vpc string `json:"vpc"`

	// vpc标识
	VpcId *string `json:"vpcId,omitempty"`

	// 微服务引擎专享版子网ID
	NetworkId string `json:"networkId"`

	// 微服务引擎专享版子网划分
	SubnetCidr string `json:"subnetCidr"`

	// 微服务引擎专享版公网地址ID，当前为null
	PublicIpId *string `json:"publicIpId,omitempty"`

	AuthCred *EngineRbacPwd `json:"auth_cred,omitempty"`

	// 微服务引擎专享版应用部署类型
	SpecType EngineCreateReqSpecType `json:"specType"`

	// 引擎附加参数
	Inputs map[string]string `json:"inputs,omitempty"`
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
	CSE_S1_SMALL2  EngineCreateReqFlavor
	CSE_S1_MEDIUM2 EngineCreateReqFlavor
	CSE_S1_LARGE2  EngineCreateReqFlavor
	CSE_S1_XLARGE2 EngineCreateReqFlavor
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
	CSE  EngineCreateReqSpecType
	CSE2 EngineCreateReqSpecType
}

func GetEngineCreateReqSpecTypeEnum() EngineCreateReqSpecTypeEnum {
	return EngineCreateReqSpecTypeEnum{
		CSE: EngineCreateReqSpecType{
			value: "CSE",
		},
		CSE2: EngineCreateReqSpecType{
			value: "CSE2",
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
