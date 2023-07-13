package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoutetableRequest Request Object
type CreateRoutetableRequest struct {
	Body *CreateRoutetableRequestBody `json:"body,omitempty"`
}

func (o CreateRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableRequest struct{}"
	}

	return strings.Join([]string{"CreateRoutetableRequest", string(data)}, " ")
}
