package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpecConvertProductResponse Response Object
type ShowSpecConvertProductResponse struct {

	// **参数解释**： 新产品ID。 **取值范围**： 不涉及。
	ProductId      *string `json:"product_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSpecConvertProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecConvertProductResponse struct{}"
	}

	return strings.Join([]string{"ShowSpecConvertProductResponse", string(data)}, " ")
}
