package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConformancePackRequest Request Object
type UpdateConformancePackRequest struct {

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`

	Body *UpdateConformancePackRequestBody `json:"body,omitempty"`
}

func (o UpdateConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConformancePackRequest struct{}"
	}

	return strings.Join([]string{"UpdateConformancePackRequest", string(data)}, " ")
}
