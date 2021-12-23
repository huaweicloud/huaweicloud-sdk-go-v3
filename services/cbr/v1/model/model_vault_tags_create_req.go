package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type VaultTagsCreateReq struct {
	Tag *Tag `json:"tag,omitempty"`
}

func (o VaultTagsCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultTagsCreateReq struct{}"
	}

	return strings.Join([]string{"VaultTagsCreateReq", string(data)}, " ")
}
