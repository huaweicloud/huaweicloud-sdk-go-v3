package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQualityEnhanceTemplateResponse Response Object
type UpdateQualityEnhanceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateQualityEnhanceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateResponse", string(data)}, " ")
}
