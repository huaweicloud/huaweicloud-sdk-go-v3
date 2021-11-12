package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListDomainsRequest struct {
	// 域名的认证状态:   * unauth - 未认证   * auth - 已认证   * invalid - 认证文件无效   * manual - 人工认证

	AuthStatus *ListDomainsRequestAuthStatus `json:"auth_status,omitempty"`
	// 分页查询，偏移量，表示从此偏移量开始查询

	Offset *int32 `json:"offset,omitempty"`
	// 分页查询，每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainsRequest", string(data)}, " ")
}

type ListDomainsRequestAuthStatus struct {
	value string
}

type ListDomainsRequestAuthStatusEnum struct {
	UNAUTH  ListDomainsRequestAuthStatus
	AUTH    ListDomainsRequestAuthStatus
	INVALID ListDomainsRequestAuthStatus
	MANUAL  ListDomainsRequestAuthStatus
}

func GetListDomainsRequestAuthStatusEnum() ListDomainsRequestAuthStatusEnum {
	return ListDomainsRequestAuthStatusEnum{
		UNAUTH: ListDomainsRequestAuthStatus{
			value: "unauth",
		},
		AUTH: ListDomainsRequestAuthStatus{
			value: "auth",
		},
		INVALID: ListDomainsRequestAuthStatus{
			value: "invalid",
		},
		MANUAL: ListDomainsRequestAuthStatus{
			value: "manual",
		},
	}
}

func (c ListDomainsRequestAuthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainsRequestAuthStatus) UnmarshalJSON(b []byte) error {
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
