package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateClassificationToEntityResponse Response Object
type AssociateClassificationToEntityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateClassificationToEntityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateClassificationToEntityResponse struct{}"
	}

	return strings.Join([]string{"AssociateClassificationToEntityResponse", string(data)}, " ")
}
