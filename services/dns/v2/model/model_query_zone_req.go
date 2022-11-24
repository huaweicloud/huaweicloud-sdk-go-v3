package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryZoneReq struct {

	// 域名名称。
	Name *string `json:"name,omitempty"`
}

func (o QueryZoneReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryZoneReq struct{}"
	}

	return strings.Join([]string{"QueryZoneReq", string(data)}, " ")
}
