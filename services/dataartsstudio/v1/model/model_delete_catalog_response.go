package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCatalogResponse Response Object
type DeleteCatalogResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCatalogResponse struct{}"
	}

	return strings.Join([]string{"DeleteCatalogResponse", string(data)}, " ")
}
