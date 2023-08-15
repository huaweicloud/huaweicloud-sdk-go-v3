package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogInfo struct {

	// 标签guid
	Guid *string `json:"guid,omitempty"`
}

func (o CatalogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogInfo struct{}"
	}

	return strings.Join([]string{"CatalogInfo", string(data)}, " ")
}
