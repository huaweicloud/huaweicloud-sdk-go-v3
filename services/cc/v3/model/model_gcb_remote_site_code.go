package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbRemoteSiteCode struct {

	// 功能说明：远端接入点的编码。
	RemoteSiteCode *string `json:"remote_site_code,omitempty"`
}

func (o GcbRemoteSiteCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbRemoteSiteCode struct{}"
	}

	return strings.Join([]string{"GcbRemoteSiteCode", string(data)}, " ")
}
