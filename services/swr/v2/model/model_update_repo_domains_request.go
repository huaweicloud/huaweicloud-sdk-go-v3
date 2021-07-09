package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateRepoDomainsRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType UpdateRepoDomainsRequestContentType `json:"Content-Type"`
	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 共享账号

	AccessDomain string `json:"access_domain"`

	Body *UpdateRepoDomainsRequestBody `json:"body,omitempty"`
}

func (o UpdateRepoDomainsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRepoDomainsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRepoDomainsRequest", string(data)}, " ")
}

type UpdateRepoDomainsRequestContentType struct {
	value string
}

type UpdateRepoDomainsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 UpdateRepoDomainsRequestContentType
	APPLICATION_JSON             UpdateRepoDomainsRequestContentType
}

func GetUpdateRepoDomainsRequestContentTypeEnum() UpdateRepoDomainsRequestContentTypeEnum {
	return UpdateRepoDomainsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: UpdateRepoDomainsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: UpdateRepoDomainsRequestContentType{
			value: "application/json",
		},
	}
}

func (c UpdateRepoDomainsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateRepoDomainsRequestContentType) UnmarshalJSON(b []byte) error {
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
