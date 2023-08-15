package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConformancePackRequest Request Object
type DeleteConformancePackRequest struct {

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`
}

func (o DeleteConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConformancePackRequest struct{}"
	}

	return strings.Join([]string{"DeleteConformancePackRequest", string(data)}, " ")
}
