package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeTrainingJobDescriptionResponse Response Object
type ChangeTrainingJobDescriptionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeTrainingJobDescriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTrainingJobDescriptionResponse struct{}"
	}

	return strings.Join([]string{"ChangeTrainingJobDescriptionResponse", string(data)}, " ")
}
