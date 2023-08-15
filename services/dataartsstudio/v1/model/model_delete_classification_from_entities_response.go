package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClassificationFromEntitiesResponse Response Object
type DeleteClassificationFromEntitiesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClassificationFromEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClassificationFromEntitiesResponse struct{}"
	}

	return strings.Join([]string{"DeleteClassificationFromEntitiesResponse", string(data)}, " ")
}
