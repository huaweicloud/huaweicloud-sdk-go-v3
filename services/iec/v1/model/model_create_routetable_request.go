package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRoutetableRequest struct {
	Body *CreateRoutetableRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableRequest struct{}"
	}

	return strings.Join([]string{"CreateRoutetableRequest", string(data)}, " ")
}
