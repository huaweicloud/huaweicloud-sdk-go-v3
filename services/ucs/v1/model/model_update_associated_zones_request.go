package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAssociatedZonesRequest struct {

	// 在联邦管理的域名访问功能中，用于更改根域名
	DnsSuffix *[]string `json:"dnsSuffix,omitempty"`
}

func (o UpdateAssociatedZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssociatedZonesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssociatedZonesRequest", string(data)}, " ")
}
