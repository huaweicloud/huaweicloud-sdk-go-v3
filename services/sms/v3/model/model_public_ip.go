package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIp 公网ip
type PublicIp struct {

	// 弹性公网IP类型，默认为5_bgp。 详细类型请参考EIP服务API文档中“查询弹性公网IP详情”部分，查看响应参数的中type字段描述
	Type string `json:"type"`

	// 带宽大小，单位：Mbit/s。 调整带宽时的最小单位会根据带宽范围不同存在差异。 小于等于300Mbit/s，默认最小单位为1Mbit/s。300Mbit/s~1000Mbit/s，默认最小单位为50Mbit/s。大于1000Mbit/s：默认最小单位为500Mbit/s。
	BandwidthSize int32 `json:"bandwidth_size"`

	// 带宽共享类型（长文本信息，非枚举数据，来源于EIP服务） 详细类型请参考EIP服务API文档中“查询弹性公网IP详情”部分，查看响应参数的中bandwidth_share_type字段描述
	BandwidthShareType *string `json:"bandwidth_share_type,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
