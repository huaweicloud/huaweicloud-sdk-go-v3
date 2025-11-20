package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CertDuration 集群证书有效期
type CertDuration struct {

	// **参数解释：** 集群证书有效时间 **约束限制：** 不涉及 **取值范围：** -1或[1,1827] > - 最小值为1天，最大值为5年，因此取值范围为1-1827（以天为单位，实际上限取决于5年内闰年的数量，例如5年内存在一个闰年则上限为1826天）； > - 若填-1则为最大值5年。  **默认取值：** 不涉及
	Duration int32 `json:"duration"`
}

func (o CertDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDuration struct{}"
	}

	return strings.Join([]string{"CertDuration", string(data)}, " ")
}
