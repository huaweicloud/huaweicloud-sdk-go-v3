package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteCodeDef 站点编码定义
type SiteCodeDef struct {
}

func (o SiteCodeDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteCodeDef struct{}"
	}

	return strings.Join([]string{"SiteCodeDef", string(data)}, " ")
}
