package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdpDetailVo 身份提供方信息
type IdpDetailVo struct {
	AddIdp *IdpInfo `json:"add_idp,omitempty"`

	UpdateIdp *IdpInfo `json:"update_idp,omitempty"`

	DelIdp *IdpInfo `json:"del_idp,omitempty"`
}

func (o IdpDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpDetailVo struct{}"
	}

	return strings.Join([]string{"IdpDetailVo", string(data)}, " ")
}
