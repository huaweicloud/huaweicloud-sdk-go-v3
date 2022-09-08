package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群证书有效期
type CertDuration struct {

	// 集群证书有效时间，单位为天，最小值为1，最大值为从当前日期起5年对应的天数，可能为1826或1827，取决于包含几个闰年的2月29日；若填-1则为最大值5年。
	Duration int32 `json:"duration"`
}

func (o CertDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDuration struct{}"
	}

	return strings.Join([]string{"CertDuration", string(data)}, " ")
}
