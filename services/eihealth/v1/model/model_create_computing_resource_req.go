package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateComputingResourceReq struct {

	// 可用区
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 规格编码
	SpecCode string `json:"spec_code"`

	// 购买数量
	Count int32 `json:"count"`

	// 额外数据盘规格编码
	DataDiskSpecCode *string `json:"data_disk_spec_code,omitempty"`

	// 额外数据盘大小
	DataDiskSize *int32 `json:"data_disk_size,omitempty"`
}

func (o CreateComputingResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingResourceReq struct{}"
	}

	return strings.Join([]string{"CreateComputingResourceReq", string(data)}, " ")
}
