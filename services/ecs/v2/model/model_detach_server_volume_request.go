/*
 * ecs
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type DetachServerVolumeRequest struct {
	ServerId   string                              `json:"server_id"`
	VolumeId   string                              `json:"volume_id"`
	DeleteFlag DetachServerVolumeRequestDeleteFlag `json:"delete_flag,omitempty"`
}

func (o DetachServerVolumeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DetachServerVolumeRequest", string(data)}, " ")
}

type DetachServerVolumeRequestDeleteFlag struct {
	value string
}

type DetachServerVolumeRequestDeleteFlagEnum struct {
	_0 DetachServerVolumeRequestDeleteFlag
	_1 DetachServerVolumeRequestDeleteFlag
}

func GetDetachServerVolumeRequestDeleteFlagEnum() DetachServerVolumeRequestDeleteFlagEnum {
	return DetachServerVolumeRequestDeleteFlagEnum{
		_0: DetachServerVolumeRequestDeleteFlag{
			value: "0",
		},
		_1: DetachServerVolumeRequestDeleteFlag{
			value: "1",
		},
	}
}

func (c DetachServerVolumeRequestDeleteFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DetachServerVolumeRequestDeleteFlag) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
