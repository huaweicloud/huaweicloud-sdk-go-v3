package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchoverResp 主备倒换响应体
type SwitchoverResp struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 更新时间，格式yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdatedAt *string `json:"updated_at,omitempty"`

	SourceDb *EndpointVo `json:"source_db,omitempty"`

	TargetDb *EndpointVo `json:"target_db,omitempty"`

	// 任务方向。 - up 入云 灾备场景时对应本云为备 - down 出云 灾备场景时对应本云为主 - non-dbs 自建
	JobDirection *SwitchoverRespJobDirection `json:"job_direction,omitempty"`

	// 目标库是否只读。
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o SwitchoverResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverResp struct{}"
	}

	return strings.Join([]string{"SwitchoverResp", string(data)}, " ")
}

type SwitchoverRespJobDirection struct {
	value string
}

type SwitchoverRespJobDirectionEnum struct {
	UP      SwitchoverRespJobDirection
	DOWN    SwitchoverRespJobDirection
	NON_DBS SwitchoverRespJobDirection
}

func GetSwitchoverRespJobDirectionEnum() SwitchoverRespJobDirectionEnum {
	return SwitchoverRespJobDirectionEnum{
		UP: SwitchoverRespJobDirection{
			value: "up",
		},
		DOWN: SwitchoverRespJobDirection{
			value: "down",
		},
		NON_DBS: SwitchoverRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c SwitchoverRespJobDirection) Value() string {
	return c.value
}

func (c SwitchoverRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchoverRespJobDirection) UnmarshalJSON(b []byte) error {
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
