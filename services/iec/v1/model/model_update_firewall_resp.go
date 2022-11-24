package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type UpdateFirewallResp struct {

	// 网络ACL ID
	Id string `json:"id"`

	// 网络ACL状态。
	Status UpdateFirewallRespStatus `json:"status"`
}

func (o UpdateFirewallResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallResp struct{}"
	}

	return strings.Join([]string{"UpdateFirewallResp", string(data)}, " ")
}

type UpdateFirewallRespStatus struct {
	value string
}

type UpdateFirewallRespStatusEnum struct {
	INACTIVE UpdateFirewallRespStatus
}

func GetUpdateFirewallRespStatusEnum() UpdateFirewallRespStatusEnum {
	return UpdateFirewallRespStatusEnum{
		INACTIVE: UpdateFirewallRespStatus{
			value: "INACTIVE",
		},
	}
}

func (c UpdateFirewallRespStatus) Value() string {
	return c.value
}

func (c UpdateFirewallRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFirewallRespStatus) UnmarshalJSON(b []byte) error {
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
