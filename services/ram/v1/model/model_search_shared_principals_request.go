package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSharedPrincipalsRequest Request Object
type SearchSharedPrincipalsRequest struct {
	Body *SearchSharedPrincipalsReqBody `json:"body,omitempty"`
}

func (o SearchSharedPrincipalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSharedPrincipalsRequest struct{}"
	}

	return strings.Join([]string{"SearchSharedPrincipalsRequest", string(data)}, " ")
}
