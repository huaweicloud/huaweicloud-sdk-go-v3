package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRandomRequest struct {

	// API版本号
	VersionId string `json:"version_id"`

	Body *GenRandomRequestBody `json:"body,omitempty"`
}

func (o CreateRandomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRandomRequest struct{}"
	}

	return strings.Join([]string{"CreateRandomRequest", string(data)}, " ")
}
