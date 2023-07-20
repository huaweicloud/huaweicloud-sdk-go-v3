package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEnterpriseRoutersRequest Request Object
type ListEnterpriseRoutersRequest struct {

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 企业路由器实例状态
	State *[]ListEnterpriseRoutersRequestState `json:"state,omitempty"`

	// 根据资源ID查询，可同时查询多个。
	Id *[]string `json:"id,omitempty"`

	// 连接对应的资源ID列表
	ResourceId *[]string `json:"resource_id,omitempty"`

	// 过滤资源是否属于当前租户；取值为true时，只查属于当前租户的资源，不包括共享资源；为false时，查询当前租户及共享给该租户的资源；
	OwnedBySelf *bool `json:"owned_by_self,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListEnterpriseRoutersRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListEnterpriseRoutersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseRoutersRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseRoutersRequest", string(data)}, " ")
}

type ListEnterpriseRoutersRequestState struct {
	value string
}

type ListEnterpriseRoutersRequestStateEnum struct {
	PENDING   ListEnterpriseRoutersRequestState
	AVAILABLE ListEnterpriseRoutersRequestState
	MODIFYING ListEnterpriseRoutersRequestState
	DELETING  ListEnterpriseRoutersRequestState
	DELETED   ListEnterpriseRoutersRequestState
	FAILED    ListEnterpriseRoutersRequestState
}

func GetListEnterpriseRoutersRequestStateEnum() ListEnterpriseRoutersRequestStateEnum {
	return ListEnterpriseRoutersRequestStateEnum{
		PENDING: ListEnterpriseRoutersRequestState{
			value: "pending",
		},
		AVAILABLE: ListEnterpriseRoutersRequestState{
			value: "available",
		},
		MODIFYING: ListEnterpriseRoutersRequestState{
			value: "modifying",
		},
		DELETING: ListEnterpriseRoutersRequestState{
			value: "deleting",
		},
		DELETED: ListEnterpriseRoutersRequestState{
			value: "deleted",
		},
		FAILED: ListEnterpriseRoutersRequestState{
			value: "failed",
		},
	}
}

func (c ListEnterpriseRoutersRequestState) Value() string {
	return c.value
}

func (c ListEnterpriseRoutersRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnterpriseRoutersRequestState) UnmarshalJSON(b []byte) error {
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

type ListEnterpriseRoutersRequestSortDir struct {
	value string
}

type ListEnterpriseRoutersRequestSortDirEnum struct {
	ASC  ListEnterpriseRoutersRequestSortDir
	DESC ListEnterpriseRoutersRequestSortDir
}

func GetListEnterpriseRoutersRequestSortDirEnum() ListEnterpriseRoutersRequestSortDirEnum {
	return ListEnterpriseRoutersRequestSortDirEnum{
		ASC: ListEnterpriseRoutersRequestSortDir{
			value: "asc",
		},
		DESC: ListEnterpriseRoutersRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListEnterpriseRoutersRequestSortDir) Value() string {
	return c.value
}

func (c ListEnterpriseRoutersRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnterpriseRoutersRequestSortDir) UnmarshalJSON(b []byte) error {
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
