package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAccessModeReq 修改接入方式请求。
type UpdateAccessModeReq struct {

	// 接入模式。 - INTERNET：互联网接入。 - DEDICATED：专线接入。 - BOTH：代表两种接入方式都支持。
	AccessMode *UpdateAccessModeReqAccessMode `json:"access_mode,omitempty"`

	// 专线接入网段列表，多个网段信息用分号隔开，列表长度不超过5。
	DedicatedCidrs *string `json:"dedicated_cidrs,omitempty"`

	ApplySharedVpcDedicatedParam *ApplySharedVpcDedicatedParam `json:"apply_shared_vpc_dedicated_param,omitempty"`
}

func (o UpdateAccessModeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessModeReq struct{}"
	}

	return strings.Join([]string{"UpdateAccessModeReq", string(data)}, " ")
}

type UpdateAccessModeReqAccessMode struct {
	value string
}

type UpdateAccessModeReqAccessModeEnum struct {
	INTERNET  UpdateAccessModeReqAccessMode
	DEDICATED UpdateAccessModeReqAccessMode
	BOTH      UpdateAccessModeReqAccessMode
}

func GetUpdateAccessModeReqAccessModeEnum() UpdateAccessModeReqAccessModeEnum {
	return UpdateAccessModeReqAccessModeEnum{
		INTERNET: UpdateAccessModeReqAccessMode{
			value: "INTERNET",
		},
		DEDICATED: UpdateAccessModeReqAccessMode{
			value: "DEDICATED",
		},
		BOTH: UpdateAccessModeReqAccessMode{
			value: "BOTH",
		},
	}
}

func (c UpdateAccessModeReqAccessMode) Value() string {
	return c.value
}

func (c UpdateAccessModeReqAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAccessModeReqAccessMode) UnmarshalJSON(b []byte) error {
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
