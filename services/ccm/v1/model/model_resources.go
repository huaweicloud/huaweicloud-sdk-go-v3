package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resources struct {

	// 证书类型:  - **CERTIFICATE_AUTHORITY**: CA证书；  - **CERTIFICATE**: 私有证书。
	Type string `json:"type"`

	// 已使用配额数。
	Used int32 `json:"used"`

	// 配额总数：   - **CERTIFICATE_AUTHORITY**: 当前系统指定100；   - **CERTIFICATE**: 当前系统指定100000。
	Quota int32 `json:"quota"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
