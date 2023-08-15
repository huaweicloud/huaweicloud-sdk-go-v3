package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNameSpaceListResponse Response Object
type ShowNameSpaceListResponse struct {
	Body           *[]NamespaceVo `json:"body,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowNameSpaceListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNameSpaceListResponse struct{}"
	}

	return strings.Join([]string{"ShowNameSpaceListResponse", string(data)}, " ")
}
