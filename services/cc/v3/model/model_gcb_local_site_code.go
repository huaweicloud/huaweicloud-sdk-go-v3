package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbLocalSiteCode struct {

	// 功能说明：本端接入点的编码。
	LocalSiteCode *string `json:"local_site_code,omitempty"`
}

func (o GcbLocalSiteCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbLocalSiteCode struct{}"
	}

	return strings.Join([]string{"GcbLocalSiteCode", string(data)}, " ")
}
