package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelValuesAomPromGetResponse Response Object
type ListLabelValuesAomPromGetResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 标签值信息。
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLabelValuesAomPromGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelValuesAomPromGetResponse struct{}"
	}

	return strings.Join([]string{"ListLabelValuesAomPromGetResponse", string(data)}, " ")
}
