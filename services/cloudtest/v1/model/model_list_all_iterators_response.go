package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllIteratorsResponse Response Object
type ListAllIteratorsResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueListTestVersionVo `json:"result,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAllIteratorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllIteratorsResponse struct{}"
	}

	return strings.Join([]string{"ListAllIteratorsResponse", string(data)}, " ")
}
