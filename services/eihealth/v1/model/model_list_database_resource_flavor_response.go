package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseResourceFlavorResponse Response Object
type ListDatabaseResourceFlavorResponse struct {

	// 个数
	Count *int32 `json:"count,omitempty"`

	// 规格列表
	Flavors        *[]DatabaseFlavorRsp `json:"flavors,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListDatabaseResourceFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseResourceFlavorResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseResourceFlavorResponse", string(data)}, " ")
}
