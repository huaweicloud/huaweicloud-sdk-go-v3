package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorRequest Request Object
type CreateIndicatorRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// workspace id
	WorkspaceId string `json:"workspace_id"`

	Body *IndicatorCreateRequest `json:"body,omitempty"`
}

func (o CreateIndicatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorRequest struct{}"
	}

	return strings.Join([]string{"CreateIndicatorRequest", string(data)}, " ")
}
