package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListNatGatewaysRequest Request Object
type ListNatGatewaysRequest struct {

	// 项目的ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 公网NAT网关实例的ID。
	Id *string `json:"id,omitempty"`

	// 企业项目ID。创建公网NAT网关实例时，关联的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 公网NAT网关实例的描述，长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 公网NAT网关实例的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 公网NAT网关实例的名字，长度限制为64。 公网NAT网关实例的名字仅支持数字、字母、_（下划线）、-（中划线）、中文
	Name *string `json:"name,omitempty"`

	// 公网NAT网关实例的状态。 取值为:  ACTIVE: 可用 PENDING_CREATE: 创建中 PENDING_UPDATE: 更新中 PENDING_DELETE: 删除中 INACTIVE: 不可用
	Status *[]ListNatGatewaysRequestStatus `json:"status,omitempty"`

	// 公网NAT网关实例的规格。 取值为： \"1\"：小型，SNAT最大连接数10000 \"2\"：中型，SNAT最大连接数50000 \"3\"：大型，SNAT最大连接数200000 \"4\"：超大型，SNAT最大连接数1000000 “5”：企业型，SNAT最大连接数10000000
	Spec *[]ListNatGatewaysRequestSpec `json:"spec,omitempty"`

	// 解冻/冻结状态。 取值范围： \"true\"：解冻 \"false\"：冻结
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 公网NAT网关下行口（DVR的下一跳）所属的network id。
	InternalNetworkId *string `json:"internal_network_id,omitempty"`

	// VPC的id。
	RouterId *string `json:"router_id,omitempty"`

	// 功能说明：每页返回的个数。 取值范围：1~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始资源ID，表示从指定资源的下一条记录开始查询。 - 若不传入marker和limit参数，查询结果返回第一页全部资源记录（默认2000条）。 - 若不传入marker参数，limit为10，查询结果返回第1~10条资源记录。 - 若marker为第10条记录的资源ID，limit为10，查询结果返回第11~20条资源记录。 - 若marker为第10条记录的资源ID，不传入limit参数，查询结果返回第11条及之后的资源记录（默认2000条）。
	Marker *string `json:"marker,omitempty"`

	// 排序使用的key
	SortKey *ListNatGatewaysRequestSortKey `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认降序desc，升序为asc
	SortDir *ListNatGatewaysRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListNatGatewaysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaysRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewaysRequest", string(data)}, " ")
}

type ListNatGatewaysRequestStatus struct {
	value string
}

type ListNatGatewaysRequestStatusEnum struct {
	ACTIVE         ListNatGatewaysRequestStatus
	PENDING_CREATE ListNatGatewaysRequestStatus
	PENDING_UPDATE ListNatGatewaysRequestStatus
	PENDING_DELETE ListNatGatewaysRequestStatus
	INACTIVE       ListNatGatewaysRequestStatus
}

func GetListNatGatewaysRequestStatusEnum() ListNatGatewaysRequestStatusEnum {
	return ListNatGatewaysRequestStatusEnum{
		ACTIVE: ListNatGatewaysRequestStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: ListNatGatewaysRequestStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: ListNatGatewaysRequestStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: ListNatGatewaysRequestStatus{
			value: "PENDING_DELETE",
		},
		INACTIVE: ListNatGatewaysRequestStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListNatGatewaysRequestStatus) Value() string {
	return c.value
}

func (c ListNatGatewaysRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaysRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListNatGatewaysRequestSpec struct {
	value string
}

type ListNatGatewaysRequestSpecEnum struct {
	E_1 ListNatGatewaysRequestSpec
	E_2 ListNatGatewaysRequestSpec
	E_3 ListNatGatewaysRequestSpec
	E_4 ListNatGatewaysRequestSpec
	E_5 ListNatGatewaysRequestSpec
}

func GetListNatGatewaysRequestSpecEnum() ListNatGatewaysRequestSpecEnum {
	return ListNatGatewaysRequestSpecEnum{
		E_1: ListNatGatewaysRequestSpec{
			value: "1",
		},
		E_2: ListNatGatewaysRequestSpec{
			value: "2",
		},
		E_3: ListNatGatewaysRequestSpec{
			value: "3",
		},
		E_4: ListNatGatewaysRequestSpec{
			value: "4",
		},
		E_5: ListNatGatewaysRequestSpec{
			value: "5",
		},
	}
}

func (c ListNatGatewaysRequestSpec) Value() string {
	return c.value
}

func (c ListNatGatewaysRequestSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaysRequestSpec) UnmarshalJSON(b []byte) error {
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

type ListNatGatewaysRequestSortKey struct {
	value string
}

type ListNatGatewaysRequestSortKeyEnum struct {
	ID         ListNatGatewaysRequestSortKey
	NAME       ListNatGatewaysRequestSortKey
	STATUS     ListNatGatewaysRequestSortKey
	CREATED_AT ListNatGatewaysRequestSortKey
}

func GetListNatGatewaysRequestSortKeyEnum() ListNatGatewaysRequestSortKeyEnum {
	return ListNatGatewaysRequestSortKeyEnum{
		ID: ListNatGatewaysRequestSortKey{
			value: "id",
		},
		NAME: ListNatGatewaysRequestSortKey{
			value: "name",
		},
		STATUS: ListNatGatewaysRequestSortKey{
			value: "status",
		},
		CREATED_AT: ListNatGatewaysRequestSortKey{
			value: "created_at",
		},
	}
}

func (c ListNatGatewaysRequestSortKey) Value() string {
	return c.value
}

func (c ListNatGatewaysRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaysRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListNatGatewaysRequestSortDir struct {
	value string
}

type ListNatGatewaysRequestSortDirEnum struct {
	DESC ListNatGatewaysRequestSortDir
	ASC  ListNatGatewaysRequestSortDir
}

func GetListNatGatewaysRequestSortDirEnum() ListNatGatewaysRequestSortDirEnum {
	return ListNatGatewaysRequestSortDirEnum{
		DESC: ListNatGatewaysRequestSortDir{
			value: "desc",
		},
		ASC: ListNatGatewaysRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListNatGatewaysRequestSortDir) Value() string {
	return c.value
}

func (c ListNatGatewaysRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewaysRequestSortDir) UnmarshalJSON(b []byte) error {
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
