package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordSetRequest Request Object
type CreateRecordSetRequest struct {

	// 项目ID
	XZoneProjectID string `json:"X-Zone-Project-ID"`

	// DNS区域ID
	XZoneID string `json:"X-Zone-ID"`

	Body *[]CreateRecordSetRequestBody `json:"body,omitempty"`
}

func (o CreateRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequest", string(data)}, " ")
}
