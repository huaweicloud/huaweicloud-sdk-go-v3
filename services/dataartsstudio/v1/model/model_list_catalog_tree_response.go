package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogTreeResponse Response Object
type ListCatalogTreeResponse struct {
	Data           *ListCatalogTreeResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListCatalogTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogTreeResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogTreeResponse", string(data)}, " ")
}
