package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsProtectModeObject struct {

	// ips防护模式id
	Id *string `json:"id,omitempty"`

	// ips防护模式，0：观察模式，1：严格模式，2：中等模式，3：宽松模式
	Mode *int32 `json:"mode,omitempty"`
}

func (o IpsProtectModeObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsProtectModeObject struct{}"
	}

	return strings.Join([]string{"IpsProtectModeObject", string(data)}, " ")
}
