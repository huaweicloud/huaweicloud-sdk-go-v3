package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UnbindApiForAcl struct {

	// API的ID
	Id *string `json:"id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// API开放状态
	Type *int32 `json:"type,omitempty"`

	// API描述
	Remark *string `json:"remark,omitempty"`

	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty"`

	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty"`

	// API发布记录编号
	PublishId *string `json:"publish_id,omitempty"`

	// 绑定的其他同类型的ACL策略名称
	AclName *string `json:"acl_name,omitempty"`

	// API的请求地址
	ReqUri *string `json:"req_uri,omitempty"`

	// API的认证方式
	AuthType *string `json:"auth_type,omitempty"`

	// API请求方法
	ReqMethod *UnbindApiForAclReqMethod `json:"req_method,omitempty"`

	// API绑定的标签，标签配额默认10条，可以联系技术调整。
	Tags *[]string `json:"tags,omitempty"`
}

func (o UnbindApiForAcl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindApiForAcl struct{}"
	}

	return strings.Join([]string{"UnbindApiForAcl", string(data)}, " ")
}

type UnbindApiForAclReqMethod struct {
	value string
}

type UnbindApiForAclReqMethodEnum struct {
	GET     UnbindApiForAclReqMethod
	POST    UnbindApiForAclReqMethod
	DELETE  UnbindApiForAclReqMethod
	PUT     UnbindApiForAclReqMethod
	PATCH   UnbindApiForAclReqMethod
	HEAD    UnbindApiForAclReqMethod
	OPTIONS UnbindApiForAclReqMethod
	ANY     UnbindApiForAclReqMethod
}

func GetUnbindApiForAclReqMethodEnum() UnbindApiForAclReqMethodEnum {
	return UnbindApiForAclReqMethodEnum{
		GET: UnbindApiForAclReqMethod{
			value: "GET",
		},
		POST: UnbindApiForAclReqMethod{
			value: "POST",
		},
		DELETE: UnbindApiForAclReqMethod{
			value: "DELETE",
		},
		PUT: UnbindApiForAclReqMethod{
			value: "PUT",
		},
		PATCH: UnbindApiForAclReqMethod{
			value: "PATCH",
		},
		HEAD: UnbindApiForAclReqMethod{
			value: "HEAD",
		},
		OPTIONS: UnbindApiForAclReqMethod{
			value: "OPTIONS",
		},
		ANY: UnbindApiForAclReqMethod{
			value: "ANY",
		},
	}
}

func (c UnbindApiForAclReqMethod) Value() string {
	return c.value
}

func (c UnbindApiForAclReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnbindApiForAclReqMethod) UnmarshalJSON(b []byte) error {
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
