package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAccessConfigRequest struct {
	Body *GetAccessConfigListRequestBody `json:"body,omitempty"`
}

func (o ListAccessConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessConfigRequest struct{}"
	}

	return strings.Join([]string{"ListAccessConfigRequest", string(data)}, " ")
}
