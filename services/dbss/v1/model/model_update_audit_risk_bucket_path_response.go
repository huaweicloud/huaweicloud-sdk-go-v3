package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditRiskBucketPathResponse Response Object
type UpdateAuditRiskBucketPathResponse struct {

	// 操作结果  - SUCCESS: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditRiskBucketPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditRiskBucketPathResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditRiskBucketPathResponse", string(data)}, " ")
}
