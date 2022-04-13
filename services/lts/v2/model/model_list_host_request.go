package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHostRequest struct {
	Body *GetHostListRequestBody `json:"body,omitempty"`
}

func (o ListHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRequest struct{}"
	}

	return strings.Join([]string{"ListHostRequest", string(data)}, " ")
}
