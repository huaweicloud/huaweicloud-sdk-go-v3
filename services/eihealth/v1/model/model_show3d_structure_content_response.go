package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Show3dStructureContentResponse Response Object
type Show3dStructureContentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o Show3dStructureContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Show3dStructureContentResponse struct{}"
	}

	return strings.Join([]string{"Show3dStructureContentResponse", string(data)}, " ")
}
