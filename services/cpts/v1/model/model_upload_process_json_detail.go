package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadProcessJsonDetail struct {

	// id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// status
	Status *int32 `json:"status,omitempty" xml:"status"`

	// cause
	Cause *string `json:"cause,omitempty" xml:"cause"`
}

func (o UploadProcessJsonDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadProcessJsonDetail struct{}"
	}

	return strings.Join([]string{"UploadProcessJsonDetail", string(data)}, " ")
}
