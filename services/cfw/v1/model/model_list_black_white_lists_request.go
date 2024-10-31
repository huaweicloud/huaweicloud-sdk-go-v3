package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBlackWhiteListsRequest Request Object
type ListBlackWhiteListsRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得
	ObjectId string `json:"object_id"`

	// 黑白名单类型4：黑名单，5：白名单
	ListType ListBlackWhiteListsRequestListType `json:"list_type"`

	// ip地址类型0：ipv4，1:ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListBlackWhiteListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlackWhiteListsRequest struct{}"
	}

	return strings.Join([]string{"ListBlackWhiteListsRequest", string(data)}, " ")
}

type ListBlackWhiteListsRequestListType struct {
	value int32
}

type ListBlackWhiteListsRequestListTypeEnum struct {
	E_4 ListBlackWhiteListsRequestListType
	E_5 ListBlackWhiteListsRequestListType
}

func GetListBlackWhiteListsRequestListTypeEnum() ListBlackWhiteListsRequestListTypeEnum {
	return ListBlackWhiteListsRequestListTypeEnum{
		E_4: ListBlackWhiteListsRequestListType{
			value: 4,
		}, E_5: ListBlackWhiteListsRequestListType{
			value: 5,
		},
	}
}

func (c ListBlackWhiteListsRequestListType) Value() int32 {
	return c.value
}

func (c ListBlackWhiteListsRequestListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBlackWhiteListsRequestListType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
