package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatabaseInstancesRequest Request Object
type ListDatabaseInstancesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。  “\\*”为系统保留字符，如果id是以“\\*”起始，表示按照\\*后面的值模糊匹配，否则，按照id精确匹配查询。不能只传入“\\*”。
	Id *string `json:"id,omitempty"`

	// 实例名称。  “\\*”为系统保留字符，如果name是以“\\*”起始，表示按照\\*后面的值模糊匹配，否则，按照name精确匹配查询。不能只传入“\\*”。
	Name *string `json:"name,omitempty"`

	// 按照实例类型查询。目前仅支持取值“Enterprise”（区分大小写），对应分布式实例（企业版）。当前支持取值\"Ha\"（区分大小写），对应主备式实例。
	Type *ListDatabaseInstancesRequestType `json:"type,omitempty"`

	// 数据库类型，区分大小写。  - GaussDB
	DatastoreType *ListDatabaseInstancesRequestDatastoreType `json:"datastore_type,omitempty"`

	// 虚拟私有云ID。  方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html)。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网的网络ID信息。  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 根据实例标签键值对进行查询。  {key}表示标签键，不可以为空或重复。最大长度127个unicode字符。key不能为空或者空字符串，不能为空格，使用之前先trim前后半角空格。不能包含+/?#&=,%特殊字符。  {value}表示标签值，可以为空。最大长度255个unicode字符，使用之前先trim 前后半角空格。不能包含+/?#&=,%特殊字符。  如果value为空，则表示any_value（查询任意value）。  如果同时使用多个标签键值对进行查询，中间使用逗号分隔开，最多包含10组。
	Tags *[]string `json:"tags,omitempty"`

	// 计费模式。  取值范围：  postPaid：后付费，即按需付费。  prePaid：预付费，即包年/包月。
	ChargeMode *ListDatabaseInstancesRequestChargeMode `json:"charge_mode,omitempty"`
}

func (o ListDatabaseInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseInstancesRequest", string(data)}, " ")
}

type ListDatabaseInstancesRequestType struct {
	value string
}

type ListDatabaseInstancesRequestTypeEnum struct {
	ENTERPRISE ListDatabaseInstancesRequestType
	HA         ListDatabaseInstancesRequestType
}

func GetListDatabaseInstancesRequestTypeEnum() ListDatabaseInstancesRequestTypeEnum {
	return ListDatabaseInstancesRequestTypeEnum{
		ENTERPRISE: ListDatabaseInstancesRequestType{
			value: "Enterprise",
		},
		HA: ListDatabaseInstancesRequestType{
			value: "Ha",
		},
	}
}

func (c ListDatabaseInstancesRequestType) Value() string {
	return c.value
}

func (c ListDatabaseInstancesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatabaseInstancesRequestType) UnmarshalJSON(b []byte) error {
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

type ListDatabaseInstancesRequestDatastoreType struct {
	value string
}

type ListDatabaseInstancesRequestDatastoreTypeEnum struct {
	GAUSS_DB ListDatabaseInstancesRequestDatastoreType
}

func GetListDatabaseInstancesRequestDatastoreTypeEnum() ListDatabaseInstancesRequestDatastoreTypeEnum {
	return ListDatabaseInstancesRequestDatastoreTypeEnum{
		GAUSS_DB: ListDatabaseInstancesRequestDatastoreType{
			value: "GaussDB",
		},
	}
}

func (c ListDatabaseInstancesRequestDatastoreType) Value() string {
	return c.value
}

func (c ListDatabaseInstancesRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatabaseInstancesRequestDatastoreType) UnmarshalJSON(b []byte) error {
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

type ListDatabaseInstancesRequestChargeMode struct {
	value string
}

type ListDatabaseInstancesRequestChargeModeEnum struct {
	POST_PAID ListDatabaseInstancesRequestChargeMode
	PRE_PAID  ListDatabaseInstancesRequestChargeMode
}

func GetListDatabaseInstancesRequestChargeModeEnum() ListDatabaseInstancesRequestChargeModeEnum {
	return ListDatabaseInstancesRequestChargeModeEnum{
		POST_PAID: ListDatabaseInstancesRequestChargeMode{
			value: "postPaid",
		},
		PRE_PAID: ListDatabaseInstancesRequestChargeMode{
			value: "prePaid",
		},
	}
}

func (c ListDatabaseInstancesRequestChargeMode) Value() string {
	return c.value
}

func (c ListDatabaseInstancesRequestChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatabaseInstancesRequestChargeMode) UnmarshalJSON(b []byte) error {
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
