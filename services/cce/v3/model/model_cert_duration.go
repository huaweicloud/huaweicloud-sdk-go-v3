package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群证书有效期
type CertDuration struct {

	// 集群证书有效时间，单位为天，最小值为1，最大值为10950(30*365，1年固定计365天，忽略闰年影响)；若填-1则为最大值30年。
	Duration int32 `json:"duration"`
}

func (o CertDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDuration struct{}"
	}

	return strings.Join([]string{"CertDuration", string(data)}, " ")
}
