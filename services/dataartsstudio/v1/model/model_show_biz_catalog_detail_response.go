package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBizCatalogDetailResponse Response Object
type ShowBizCatalogDetailResponse struct {
	Data           *CreateCatalogResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowBizCatalogDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBizCatalogDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowBizCatalogDetailResponse", string(data)}, " ")
}
