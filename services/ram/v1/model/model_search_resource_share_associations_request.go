package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceShareAssociationsRequest Request Object
type SearchResourceShareAssociationsRequest struct {
	Body *SearchResourceShareAssociationsReqBody `json:"body,omitempty"`
}

func (o SearchResourceShareAssociationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareAssociationsRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceShareAssociationsRequest", string(data)}, " ")
}
