package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapFlavorResponse Response Object
type ListHtapFlavorResponse struct {

	// 规格信息。
	Flavors        *[]HtapFlavorInfoFlavors `json:"flavors,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListHtapFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapFlavorResponse struct{}"
	}

	return strings.Join([]string{"ListHtapFlavorResponse", string(data)}, " ")
}
