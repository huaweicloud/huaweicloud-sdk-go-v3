package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBlackWhiteListsRequest Request Object
type ListBlackWhiteListsRequest struct {

	// 互联网边界防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，type为0的为互联网边界防护对象id。
	ObjectId string `json:"object_id"`

	// 黑白名单类型4：黑名单，5：白名单
	ListType ListBlackWhiteListsRequestListType `json:"list_type"`

	// IP地址类型0：ipv4,1:ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
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
