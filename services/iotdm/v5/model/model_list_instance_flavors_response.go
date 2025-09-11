package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceFlavorsResponse Response Object
type ListInstanceFlavorsResponse struct {
	Page *Page `json:"page,omitempty"`

	// **参数说明**：设备接入实例规格的详情列表。
	Flavors        *[]InstanceFlavor `json:"flavors,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListInstanceFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceFlavorsResponse", string(data)}, " ")
}
