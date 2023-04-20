package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CDN状态码缓存时间
type ErrorCodeCache struct {

	// 允许配置的错误码: 400, 403, 404, 405, 414, 500, 501, 502, 503, 504
	Code *ErrorCodeCacheCode `json:"code,omitempty"`

	// 错误码缓存时间，单位为秒，范围0-31,536,000(一年默认为365天)
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o ErrorCodeCache) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCodeCache struct{}"
	}

	return strings.Join([]string{"ErrorCodeCache", string(data)}, " ")
}

type ErrorCodeCacheCode struct {
	value int32
}

type ErrorCodeCacheCodeEnum struct {
	E_400 ErrorCodeCacheCode
	E_403 ErrorCodeCacheCode
	E_404 ErrorCodeCacheCode
	E_405 ErrorCodeCacheCode
	E_414 ErrorCodeCacheCode
	E_500 ErrorCodeCacheCode
	E_501 ErrorCodeCacheCode
	E_502 ErrorCodeCacheCode
	E_503 ErrorCodeCacheCode
	E_504 ErrorCodeCacheCode
}

func GetErrorCodeCacheCodeEnum() ErrorCodeCacheCodeEnum {
	return ErrorCodeCacheCodeEnum{
		E_400: ErrorCodeCacheCode{
			value: 400,
		}, E_403: ErrorCodeCacheCode{
			value: 403,
		}, E_404: ErrorCodeCacheCode{
			value: 404,
		}, E_405: ErrorCodeCacheCode{
			value: 405,
		}, E_414: ErrorCodeCacheCode{
			value: 414,
		}, E_500: ErrorCodeCacheCode{
			value: 500,
		}, E_501: ErrorCodeCacheCode{
			value: 501,
		}, E_502: ErrorCodeCacheCode{
			value: 502,
		}, E_503: ErrorCodeCacheCode{
			value: 503,
		}, E_504: ErrorCodeCacheCode{
			value: 504,
		},
	}
}

func (c ErrorCodeCacheCode) Value() int32 {
	return c.value
}

func (c ErrorCodeCacheCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ErrorCodeCacheCode) UnmarshalJSON(b []byte) error {
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
