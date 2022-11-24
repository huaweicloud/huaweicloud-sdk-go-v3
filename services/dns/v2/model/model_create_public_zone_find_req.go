package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePublicZoneFindReq struct {

	// 找回域名名称
	ZoneName *string `json:"zone_name,omitempty"`
}

func (o CreatePublicZoneFindReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneFindReq struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneFindReq", string(data)}, " ")
}
