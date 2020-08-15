/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type SubnetResult struct {
	// uuid形式的一个资源标识。
	Id string `json:"id"`
	// 功能说明：子网的状态。   取值范围：ACTIVE,UNKNOWN,ERROR       ACTIVE表示子网已挂载到ROUTER上       UNKNOWN表示子网还未挂载到ROUTER上       ERROR表示子网状态故障
	Status SubnetResultStatus `json:"status"`
}

func (o SubnetResult) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SubnetResult", string(data)}, " ")
}

type SubnetResultStatus struct {
	value string
}

type SubnetResultStatusEnum struct {
	ACTIVE  SubnetResultStatus
	UNKNOWN SubnetResultStatus
	ERROR   SubnetResultStatus
}

func GetSubnetResultStatusEnum() SubnetResultStatusEnum {
	return SubnetResultStatusEnum{
		ACTIVE: SubnetResultStatus{
			value: "ACTIVE",
		},
		UNKNOWN: SubnetResultStatus{
			value: "UNKNOWN",
		},
		ERROR: SubnetResultStatus{
			value: "ERROR",
		},
	}
}

func (c SubnetResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SubnetResultStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
