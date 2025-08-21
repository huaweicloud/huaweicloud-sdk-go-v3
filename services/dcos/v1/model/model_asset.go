package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Asset struct {

	// 资产sn
	AssetSn *string `json:"asset_sn,omitempty"`

	// 资产名称
	Name *string `json:"name,omitempty"`

	// 资产类型
	Model *string `json:"model,omitempty"`

	// 资产大类
	AssetType *string `json:"asset_type,omitempty"`

	// 资产大类-中文
	MaxTypeCnName *string `json:"max_type_cn_name,omitempty"`

	// 资产大类-英文
	MaxTypeEnName *string `json:"max_type_en_name,omitempty"`

	// 资产状态
	AssetStatus *string `json:"asset_status,omitempty"`

	// dc名称
	DcName *string `json:"dc_name,omitempty"`

	// dc编码
	DcCode *string `json:"dc_code,omitempty"`

	// 房间名称
	RoomName *string `json:"room_name,omitempty"`

	// 机房编码
	RoomCode *string `json:"room_code,omitempty"`

	// 机柜编码
	RackId *string `json:"rack_id,omitempty"`

	// 首次入库时间
	FirstInboundTime *string `json:"first_inbound_time,omitempty"`

	// 出库时间
	OutboundTime *string `json:"outbound_time,omitempty"`

	// 入库时间
	InboundTime *string `json:"inbound_time,omitempty"`

	// 支架sn
	MountSn *string `json:"mount_sn,omitempty"`
}

func (o Asset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Asset struct{}"
	}

	return strings.Join([]string{"Asset", string(data)}, " ")
}
