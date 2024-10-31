package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type AclBindApiInfo struct {

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// API名称
	ApiName *string `json:"api_name,omitempty"`

	// API类型
	ApiType *int64 `json:"api_type,omitempty"`

	// API的描述信息
	ApiRemark *string `json:"api_remark,omitempty"`

	// 生效的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// 生效的环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 绑定关系编号
	BindId *string `json:"bind_id,omitempty"`

	// API分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 绑定时间
	BindTime *sdktime.SdkTime `json:"bind_time,omitempty"`

	// API发布记录编号
	PublishId *string `json:"publish_id,omitempty"`

	// API请求方法
	ReqMethod *AclBindApiInfoReqMethod `json:"req_method,omitempty"`

	// API绑定的标签，标签配额默认10条，可以联系技术调整。
	Tags *[]string `json:"tags,omitempty"`
}

func (o AclBindApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBindApiInfo struct{}"
	}

	return strings.Join([]string{"AclBindApiInfo", string(data)}, " ")
}

type AclBindApiInfoReqMethod struct {
	value string
}

type AclBindApiInfoReqMethodEnum struct {
	GET     AclBindApiInfoReqMethod
	POST    AclBindApiInfoReqMethod
	DELETE  AclBindApiInfoReqMethod
	PUT     AclBindApiInfoReqMethod
	PATCH   AclBindApiInfoReqMethod
	HEAD    AclBindApiInfoReqMethod
	OPTIONS AclBindApiInfoReqMethod
	ANY     AclBindApiInfoReqMethod
}

func GetAclBindApiInfoReqMethodEnum() AclBindApiInfoReqMethodEnum {
	return AclBindApiInfoReqMethodEnum{
		GET: AclBindApiInfoReqMethod{
			value: "GET",
		},
		POST: AclBindApiInfoReqMethod{
			value: "POST",
		},
		DELETE: AclBindApiInfoReqMethod{
			value: "DELETE",
		},
		PUT: AclBindApiInfoReqMethod{
			value: "PUT",
		},
		PATCH: AclBindApiInfoReqMethod{
			value: "PATCH",
		},
		HEAD: AclBindApiInfoReqMethod{
			value: "HEAD",
		},
		OPTIONS: AclBindApiInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: AclBindApiInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c AclBindApiInfoReqMethod) Value() string {
	return c.value
}

func (c AclBindApiInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AclBindApiInfoReqMethod) UnmarshalJSON(b []byte) error {
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
