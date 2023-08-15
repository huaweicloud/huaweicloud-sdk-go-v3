package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCatalogResponse Response Object
type ChangeCatalogResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ChangeCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCatalogResponse struct{}"
	}

	return strings.Join([]string{"ChangeCatalogResponse", string(data)}, " ")
}
