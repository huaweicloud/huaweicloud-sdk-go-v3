package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogueResponse Response Object
type CreateCatalogueResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Data *CatalogueDetailInfo `json:"data,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCatalogueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogueResponse struct{}"
	}

	return strings.Join([]string{"CreateCatalogueResponse", string(data)}, " ")
}
