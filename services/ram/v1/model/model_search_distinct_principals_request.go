package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchDistinctPrincipalsRequest struct {
	Body *SearchDistinctSharedPrincipalsReqBody `json:"body,omitempty"`
}

func (o SearchDistinctPrincipalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDistinctPrincipalsRequest struct{}"
	}

	return strings.Join([]string{"SearchDistinctPrincipalsRequest", string(data)}, " ")
}
