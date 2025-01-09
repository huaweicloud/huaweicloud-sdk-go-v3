package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppCatalogsResponse Response Object
type ListAppCatalogsResponse struct {

	// 应用分类信息。
	Items          *[]Catalog `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListAppCatalogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppCatalogsResponse struct{}"
	}

	return strings.Join([]string{"ListAppCatalogsResponse", string(data)}, " ")
}
