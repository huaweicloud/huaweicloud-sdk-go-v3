package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateClassificationToEntitiesResponse Response Object
type BatchAssociateClassificationToEntitiesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAssociateClassificationToEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateClassificationToEntitiesResponse struct{}"
	}

	return strings.Join([]string{"BatchAssociateClassificationToEntitiesResponse", string(data)}, " ")
}
