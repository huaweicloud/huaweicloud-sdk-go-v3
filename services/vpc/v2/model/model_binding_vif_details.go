package model

import (
	"encoding/json"

	"strings"
)

//
type BindingVifDetails struct {
	// 功能说明：取值为true，表示是虚拟机的主网卡。

	PrimaryInterface *bool `json:"primary_interface,omitempty"`
}

func (o BindingVifDetails) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BindingVifDetails struct{}"
	}

	return strings.Join([]string{"BindingVifDetails", string(data)}, " ")
}
