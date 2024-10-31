package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsProtectModeObject struct {

	// ips防护模式id，此处为防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得
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
