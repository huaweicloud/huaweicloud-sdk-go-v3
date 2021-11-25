package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLabelsAomPromGetResponse struct {
	Status *string `json:"status,omitempty"`

	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLabelsAomPromGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelsAomPromGetResponse struct{}"
	}

	return strings.Join([]string{"ListLabelsAomPromGetResponse", string(data)}, " ")
}
