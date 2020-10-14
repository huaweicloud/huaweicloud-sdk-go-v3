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
type ShowUpBandwidthV2Request struct {
	PublishDomains []string                          `json:"publish_domains"`
	App            *string                           `json:"app,omitempty"`
	Stream         *string                           `json:"stream,omitempty"`
	Region         *[]string                         `json:"region,omitempty"`
	Isp            *[]string                         `json:"isp,omitempty"`
	Interval       *ShowUpBandwidthV2RequestInterval `json:"interval,omitempty"`
	StartTime      *string                           `json:"start_time,omitempty"`
	EndTime        *string                           `json:"end_time,omitempty"`
}

func (o ShowUpBandwidthV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUpBandwidthV2Request", string(data)}, " ")
}

type ShowUpBandwidthV2RequestInterval struct {
	value int32
}

type ShowUpBandwidthV2RequestIntervalEnum struct {
	E_300   ShowUpBandwidthV2RequestInterval
	E_3600  ShowUpBandwidthV2RequestInterval
	E_86400 ShowUpBandwidthV2RequestInterval
}

func GetShowUpBandwidthV2RequestIntervalEnum() ShowUpBandwidthV2RequestIntervalEnum {
	return ShowUpBandwidthV2RequestIntervalEnum{
		E_300: ShowUpBandwidthV2RequestInterval{
			value: 300,
		}, E_3600: ShowUpBandwidthV2RequestInterval{
			value: 3600,
		}, E_86400: ShowUpBandwidthV2RequestInterval{
			value: 86400,
		},
	}
}

func (c ShowUpBandwidthV2RequestInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowUpBandwidthV2RequestInterval) UnmarshalJSON(b []byte) error {
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
