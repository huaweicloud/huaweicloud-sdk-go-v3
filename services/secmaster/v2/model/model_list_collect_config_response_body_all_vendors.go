package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCollectConfigResponseBodyAllVendors struct {

	// 云厂商ID
	CloudVendor *string `json:"cloud_vendor,omitempty"`

	// 所有的云产品
	CsvcList *[]ListCollectConfigResponseBodyCsvcList `json:"csvc_list,omitempty"`
}

func (o ListCollectConfigResponseBodyAllVendors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyAllVendors struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyAllVendors", string(data)}, " ")
}
