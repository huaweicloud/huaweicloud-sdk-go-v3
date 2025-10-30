package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSaleCyclesRequest Request Object
type ListSaleCyclesRequest struct {

	// 地区编码
	ZoneCode *string `json:"zone_code,omitempty"`
}

func (o ListSaleCyclesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSaleCyclesRequest struct{}"
	}

	return strings.Join([]string{"ListSaleCyclesRequest", string(data)}, " ")
}
