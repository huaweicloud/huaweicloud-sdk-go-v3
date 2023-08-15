package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSubjectsRequest Request Object
type ChangeSubjectsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *CatalogLevelVoList `json:"body,omitempty"`
}

func (o ChangeSubjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSubjectsRequest struct{}"
	}

	return strings.Join([]string{"ChangeSubjectsRequest", string(data)}, " ")
}
