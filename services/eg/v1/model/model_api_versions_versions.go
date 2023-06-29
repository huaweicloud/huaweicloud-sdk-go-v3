package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiVersionsVersions struct {

	// 版本号，如v1
	Id *string `json:"id,omitempty"`

	// url地址
	Links *[]ApiVersionsLinks `json:"links,omitempty"`

	// 若该版本API支持微版本，则为支持的最大微版本号，如果不支持微版本，则为空
	Version *string `json:"version,omitempty"`

	// 若该版本API支持微版本，则为支持的最小微版本号，如果不支持微版本，则为空
	MinVersion *string `json:"min_version,omitempty"`

	// 版本状态，支持CURRENT：推荐版本；SUPPORTED：老版本，仍支持使用；DEPRECATED：废弃版本，后续会删除
	Status *ApiVersionsVersionsStatus `json:"status,omitempty"`

	// 版本发布UTC时间
	Updated *string `json:"updated,omitempty"`
}

func (o ApiVersionsVersions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionsVersions struct{}"
	}

	return strings.Join([]string{"ApiVersionsVersions", string(data)}, " ")
}

type ApiVersionsVersionsStatus struct {
	value string
}

type ApiVersionsVersionsStatusEnum struct {
	CURRENT    ApiVersionsVersionsStatus
	SUPPORTED  ApiVersionsVersionsStatus
	DEPRECATED ApiVersionsVersionsStatus
}

func GetApiVersionsVersionsStatusEnum() ApiVersionsVersionsStatusEnum {
	return ApiVersionsVersionsStatusEnum{
		CURRENT: ApiVersionsVersionsStatus{
			value: "CURRENT",
		},
		SUPPORTED: ApiVersionsVersionsStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: ApiVersionsVersionsStatus{
			value: "DEPRECATED",
		},
	}
}

func (c ApiVersionsVersionsStatus) Value() string {
	return c.value
}

func (c ApiVersionsVersionsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiVersionsVersionsStatus) UnmarshalJSON(b []byte) error {
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
