package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorCreateRequest 创建情报请求参数
type IndicatorCreateRequest struct {
	DataObject *CreateIndicatorDetail `json:"data_object"`
}

func (o IndicatorCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorCreateRequest struct{}"
	}

	return strings.Join([]string{"IndicatorCreateRequest", string(data)}, " ")
}
