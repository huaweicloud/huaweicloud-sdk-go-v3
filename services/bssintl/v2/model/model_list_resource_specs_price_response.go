package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceSpecsPriceResponse Response Object
type ListResourceSpecsPriceResponse struct {
	PageInfo *ResourceSpecsPricePageInfo `json:"page_info,omitempty"`

	// 云服务区编码
	RegionCode *string `json:"region_code,omitempty"`

	// 资源规格列表
	ResourceSpecInfos *[]ResourceSpecInfo `json:"resource_spec_infos,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListResourceSpecsPriceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSpecsPriceResponse struct{}"
	}

	return strings.Join([]string{"ListResourceSpecsPriceResponse", string(data)}, " ")
}
