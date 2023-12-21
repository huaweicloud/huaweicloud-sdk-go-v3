package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateByConditionVoPersistableModelUpdateDto struct {
	Condition *QueryRequestVo `json:"condition,omitempty"`

	UpdateDTO *PersistableModelUpdateDto `json:"updateDTO,omitempty"`
}

func (o UpdateByConditionVoPersistableModelUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateByConditionVoPersistableModelUpdateDto struct{}"
	}

	return strings.Join([]string{"UpdateByConditionVoPersistableModelUpdateDto", string(data)}, " ")
}
