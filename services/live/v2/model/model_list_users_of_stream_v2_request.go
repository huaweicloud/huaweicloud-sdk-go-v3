/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListUsersOfStreamV2Request struct {
	PlayDomain string                              `json:"play_domain"`
	App        *string                             `json:"app,omitempty"`
	Stream     *string                             `json:"stream,omitempty"`
	Isp        *[]string                           `json:"isp,omitempty"`
	Region     *[]string                           `json:"region,omitempty"`
	Interval   *ListUsersOfStreamV2RequestInterval `json:"interval,omitempty"`
	StartTime  *string                             `json:"start_time,omitempty"`
	EndTime    *string                             `json:"end_time,omitempty"`
}

func (o ListUsersOfStreamV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListUsersOfStreamV2Request", string(data)}, " ")
}

type ListUsersOfStreamV2RequestInterval struct {
	value int32
}

type ListUsersOfStreamV2RequestIntervalEnum struct {
	E_60  ListUsersOfStreamV2RequestInterval
	E_300 ListUsersOfStreamV2RequestInterval
}

func GetListUsersOfStreamV2RequestIntervalEnum() ListUsersOfStreamV2RequestIntervalEnum {
	return ListUsersOfStreamV2RequestIntervalEnum{
		E_60: ListUsersOfStreamV2RequestInterval{
			value: 60,
		}, E_300: ListUsersOfStreamV2RequestInterval{
			value: 300,
		},
	}
}

func (c ListUsersOfStreamV2RequestInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListUsersOfStreamV2RequestInterval) UnmarshalJSON(b []byte) error {
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
