package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventTopRiskResponse Response Object
type ListEventTopRiskResponse struct {

	// TOP5事件类型列表
	DataList       *[]EventTopRiskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListEventTopRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventTopRiskResponse struct{}"
	}

	return strings.Join([]string{"ListEventTopRiskResponse", string(data)}, " ")
}
