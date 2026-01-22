package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiVirusVo struct {

	// 反病毒开关状态，0表示关闭，1表示开启
	AntiVirusStatus *int32 `json:"anti_virus_status,omitempty"`

	// 防护对象id
	Id *string `json:"id,omitempty"`

	// 防护对象名称
	Name *string `json:"name,omitempty"`
}

func (o AntiVirusVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusVo struct{}"
	}

	return strings.Join([]string{"AntiVirusVo", string(data)}, " ")
}
