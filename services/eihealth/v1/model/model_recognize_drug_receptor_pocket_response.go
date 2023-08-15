package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeDrugReceptorPocketResponse Response Object
type RecognizeDrugReceptorPocketResponse struct {
	Mode *RecognizeReceptorPocketMode `json:"mode,omitempty"`

	// 口袋列表
	Pockets        *[]BoundingBoxDto `json:"pockets,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RecognizeDrugReceptorPocketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeDrugReceptorPocketResponse struct{}"
	}

	return strings.Join([]string{"RecognizeDrugReceptorPocketResponse", string(data)}, " ")
}
