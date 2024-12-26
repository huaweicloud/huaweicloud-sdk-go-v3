package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchEipStatusDto struct {

	// 防火墙id
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 是否开启新增eip自动防护，1；是，0：否
	Status *int32 `json:"status,omitempty"`
}

func (o SwitchEipStatusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchEipStatusDto struct{}"
	}

	return strings.Join([]string{"SwitchEipStatusDto", string(data)}, " ")
}
