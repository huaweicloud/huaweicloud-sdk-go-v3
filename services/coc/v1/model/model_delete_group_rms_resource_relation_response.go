package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupRmsResourceRelationResponse Response Object
type DeleteGroupRmsResourceRelationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGroupRmsResourceRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupRmsResourceRelationResponse struct{}"
	}

	return strings.Join([]string{"DeleteGroupRmsResourceRelationResponse", string(data)}, " ")
}
