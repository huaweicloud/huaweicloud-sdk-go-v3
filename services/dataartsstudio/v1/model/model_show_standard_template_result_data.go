package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardTemplateResultData 接口返回的数据。
type ShowStandardTemplateResultData struct {
	Value *ShowStandardTemplateResultDataValue `json:"value,omitempty"`
}

func (o ShowStandardTemplateResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardTemplateResultData struct{}"
	}

	return strings.Join([]string{"ShowStandardTemplateResultData", string(data)}, " ")
}
