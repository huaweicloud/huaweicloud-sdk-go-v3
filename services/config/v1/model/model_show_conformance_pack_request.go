package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConformancePackRequest Request Object
type ShowConformancePackRequest struct {

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`
}

func (o ShowConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConformancePackRequest struct{}"
	}

	return strings.Join([]string{"ShowConformancePackRequest", string(data)}, " ")
}
