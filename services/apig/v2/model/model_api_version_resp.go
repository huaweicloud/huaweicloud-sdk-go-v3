package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiVersionResp struct {

	// API历史版本的ID
	VersionId *string `json:"version_id,omitempty"`

	// API的版本号
	VersionNo *string `json:"version_no,omitempty"`

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// 发布的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// 发布的环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 发布描述
	Remark *string `json:"remark,omitempty"`

	// 发布时间
	PublishTime *sdktime.SdkTime `json:"publish_time,omitempty"`

	// 版本状态 - 1：当前生效中的版本 - 2：未生效的版本
	Status *ApiVersionRespStatus `json:"status,omitempty"`
}

func (o ApiVersionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionResp struct{}"
	}

	return strings.Join([]string{"ApiVersionResp", string(data)}, " ")
}

type ApiVersionRespStatus struct {
	value int32
}

type ApiVersionRespStatusEnum struct {
	E_1 ApiVersionRespStatus
	E_2 ApiVersionRespStatus
}

func GetApiVersionRespStatusEnum() ApiVersionRespStatusEnum {
	return ApiVersionRespStatusEnum{
		E_1: ApiVersionRespStatus{
			value: 1,
		}, E_2: ApiVersionRespStatus{
			value: 2,
		},
	}
}

func (c ApiVersionRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiVersionRespStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
