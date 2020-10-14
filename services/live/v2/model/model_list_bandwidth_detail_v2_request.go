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
type ListBandwidthDetailV2Request struct {
	PlayDomains []string                              `json:"play_domains"`
	App         *string                               `json:"app,omitempty"`
	Stream      *string                               `json:"stream,omitempty"`
	Region      *[]string                             `json:"region,omitempty"`
	Isp         *[]string                             `json:"isp,omitempty"`
	Interval    *ListBandwidthDetailV2RequestInterval `json:"interval,omitempty"`
	StartTime   *string                               `json:"start_time,omitempty"`
	EndTime     *string                               `json:"end_time,omitempty"`
}

func (o ListBandwidthDetailV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBandwidthDetailV2Request", string(data)}, " ")
}

type ListBandwidthDetailV2RequestInterval struct {
	value int32
}

type ListBandwidthDetailV2RequestIntervalEnum struct {
	E_300   ListBandwidthDetailV2RequestInterval
	E_3600  ListBandwidthDetailV2RequestInterval
	E_86400 ListBandwidthDetailV2RequestInterval
}

func GetListBandwidthDetailV2RequestIntervalEnum() ListBandwidthDetailV2RequestIntervalEnum {
	return ListBandwidthDetailV2RequestIntervalEnum{
		E_300: ListBandwidthDetailV2RequestInterval{
			value: 300,
		}, E_3600: ListBandwidthDetailV2RequestInterval{
			value: 3600,
		}, E_86400: ListBandwidthDetailV2RequestInterval{
			value: 86400,
		},
	}
}

func (c ListBandwidthDetailV2RequestInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListBandwidthDetailV2RequestInterval) UnmarshalJSON(b []byte) error {
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
