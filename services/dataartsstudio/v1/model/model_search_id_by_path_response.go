package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchIdByPathResponse Response Object
type SearchIdByPathResponse struct {

	// 目录编号
	CatalogId      *string `json:"catalog_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchIdByPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchIdByPathResponse struct{}"
	}

	return strings.Join([]string{"SearchIdByPathResponse", string(data)}, " ")
}
