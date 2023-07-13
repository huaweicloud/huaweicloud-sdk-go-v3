package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostDiskResponse Response Object
type ListHostDiskResponse struct {
	Body           *[]DiskResp `json:"body,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListHostDiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostDiskResponse struct{}"
	}

	return strings.Join([]string{"ListHostDiskResponse", string(data)}, " ")
}
