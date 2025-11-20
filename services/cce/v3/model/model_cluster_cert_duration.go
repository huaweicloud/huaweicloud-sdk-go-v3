package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterCertDuration 集群证书有效期
type ClusterCertDuration struct {

	// **参数解释：** 集群证书有效时间。 **约束限制：** duration和expire_at参数至少需要指定一个，若同时指定则以expire_at参数为准。 **取值范围：** 最小值为1天，最大值为5年，因此取值范围为1-1827（以天为单位，实际上限取决于5年内闰年的数量，例如5年内存在一个闰年则上限为1826天）；若填-1则为最大值5年 **默认取值：** 不涉及
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释：** 集群证书到期时间。 **约束限制：** duration和expire_at参数至少需要指定一个，若同时指定则以expire_at参数为准。 **取值范围：** 证书到期时间须在当前时间后15分钟至5年之间，参数格式为：2025-01-01 16:00:00 +0000 UTC。 **默认取值：** 不涉及
	ExpireAt *string `json:"expire_at,omitempty"`
}

func (o ClusterCertDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCertDuration struct{}"
	}

	return strings.Join([]string{"ClusterCertDuration", string(data)}, " ")
}
