package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListVpcAttachmentsRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 连接状态:pending|available|modifying|deleting|deleted|failed|pending_acceptance|rejected|initiating_request
	State *[]ListVpcAttachmentsRequestState `json:"state,omitempty"`

	// 根据资源ID查询，可同时查询多个。
	Id *[]string `json:"id,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListVpcAttachmentsRequestSortDir `json:"sort_dir,omitempty"`

	// VPC id
	VpcId *[]string `json:"vpc_id,omitempty"`
}

func (o ListVpcAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcAttachmentsRequest", string(data)}, " ")
}

type ListVpcAttachmentsRequestState struct {
	value string
}

type ListVpcAttachmentsRequestStateEnum struct {
	PENDING            ListVpcAttachmentsRequestState
	AVAILABLE          ListVpcAttachmentsRequestState
	MODIFYING          ListVpcAttachmentsRequestState
	DELETING           ListVpcAttachmentsRequestState
	DELETED            ListVpcAttachmentsRequestState
	FAILED             ListVpcAttachmentsRequestState
	INITIATING_REQUEST ListVpcAttachmentsRequestState
	REJECTED           ListVpcAttachmentsRequestState
	PENDING_ACCEPTANCE ListVpcAttachmentsRequestState
}

func GetListVpcAttachmentsRequestStateEnum() ListVpcAttachmentsRequestStateEnum {
	return ListVpcAttachmentsRequestStateEnum{
		PENDING: ListVpcAttachmentsRequestState{
			value: "pending",
		},
		AVAILABLE: ListVpcAttachmentsRequestState{
			value: "available",
		},
		MODIFYING: ListVpcAttachmentsRequestState{
			value: "modifying",
		},
		DELETING: ListVpcAttachmentsRequestState{
			value: "deleting",
		},
		DELETED: ListVpcAttachmentsRequestState{
			value: "deleted",
		},
		FAILED: ListVpcAttachmentsRequestState{
			value: "failed",
		},
		INITIATING_REQUEST: ListVpcAttachmentsRequestState{
			value: "initiating_request",
		},
		REJECTED: ListVpcAttachmentsRequestState{
			value: "rejected",
		},
		PENDING_ACCEPTANCE: ListVpcAttachmentsRequestState{
			value: "pending_acceptance",
		},
	}
}

func (c ListVpcAttachmentsRequestState) Value() string {
	return c.value
}

func (c ListVpcAttachmentsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcAttachmentsRequestState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListVpcAttachmentsRequestSortDir struct {
	value string
}

type ListVpcAttachmentsRequestSortDirEnum struct {
	ASC  ListVpcAttachmentsRequestSortDir
	DESC ListVpcAttachmentsRequestSortDir
}

func GetListVpcAttachmentsRequestSortDirEnum() ListVpcAttachmentsRequestSortDirEnum {
	return ListVpcAttachmentsRequestSortDirEnum{
		ASC: ListVpcAttachmentsRequestSortDir{
			value: "asc",
		},
		DESC: ListVpcAttachmentsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListVpcAttachmentsRequestSortDir) Value() string {
	return c.value
}

func (c ListVpcAttachmentsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcAttachmentsRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
