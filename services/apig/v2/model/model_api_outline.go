package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiOutline struct {

	// API的认证方式。 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *ApiOutlineAuthType `json:"auth_type,omitempty"`

	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 发布记录的编号
	PublishId *string `json:"publish_id,omitempty"`

	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API描述
	Remark *string `json:"remark,omitempty"`

	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API的请求地址
	ReqUri *string `json:"req_uri,omitempty"`

	// API绑定的标签，标签配额默认10条，可以联系技术调整。
	Tags *[]string `json:"tags,omitempty"`
}

func (o ApiOutline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiOutline struct{}"
	}

	return strings.Join([]string{"ApiOutline", string(data)}, " ")
}

type ApiOutlineAuthType struct {
	value string
}

type ApiOutlineAuthTypeEnum struct {
	NONE       ApiOutlineAuthType
	APP        ApiOutlineAuthType
	IAM        ApiOutlineAuthType
	AUTHORIZER ApiOutlineAuthType
}

func GetApiOutlineAuthTypeEnum() ApiOutlineAuthTypeEnum {
	return ApiOutlineAuthTypeEnum{
		NONE: ApiOutlineAuthType{
			value: "NONE",
		},
		APP: ApiOutlineAuthType{
			value: "APP",
		},
		IAM: ApiOutlineAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiOutlineAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiOutlineAuthType) Value() string {
	return c.value
}

func (c ApiOutlineAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiOutlineAuthType) UnmarshalJSON(b []byte) error {
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
