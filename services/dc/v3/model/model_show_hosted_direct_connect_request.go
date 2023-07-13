package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHostedDirectConnectRequest Request Object
type ShowHostedDirectConnectRequest struct {

	// 托管专线连接ID。
	HostedConnectId string `json:"hosted_connect_id"`

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 返回结果按照升序(asc)或降序(desc)排列，默认为asc
	SortDir *[]ShowHostedDirectConnectRequestSortDir `json:"sort_dir,omitempty"`

	// 排序字段。
	SortKey *string `json:"sort_key,omitempty"`

	// 根椐运营专线ID过滤托管专线列表
	HostingId *[]string `json:"hosting_id,omitempty"`
}

func (o ShowHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"ShowHostedDirectConnectRequest", string(data)}, " ")
}

type ShowHostedDirectConnectRequestSortDir struct {
	value string
}

type ShowHostedDirectConnectRequestSortDirEnum struct {
	ASC  ShowHostedDirectConnectRequestSortDir
	DESC ShowHostedDirectConnectRequestSortDir
}

func GetShowHostedDirectConnectRequestSortDirEnum() ShowHostedDirectConnectRequestSortDirEnum {
	return ShowHostedDirectConnectRequestSortDirEnum{
		ASC: ShowHostedDirectConnectRequestSortDir{
			value: "asc",
		},
		DESC: ShowHostedDirectConnectRequestSortDir{
			value: "desc",
		},
	}
}

func (c ShowHostedDirectConnectRequestSortDir) Value() string {
	return c.value
}

func (c ShowHostedDirectConnectRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHostedDirectConnectRequestSortDir) UnmarshalJSON(b []byte) error {
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
