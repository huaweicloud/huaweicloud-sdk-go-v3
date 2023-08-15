package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseWaterMarkRequest Request Object
type CreateDatabaseWaterMarkRequest struct {
	Body *EmbeddedDatabaseWatermark `json:"body,omitempty"`
}

func (o CreateDatabaseWaterMarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseWaterMarkRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseWaterMarkRequest", string(data)}, " ")
}
