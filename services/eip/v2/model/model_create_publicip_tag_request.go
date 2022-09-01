package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePublicipTagRequest struct {

	// 资源ID
	PublicipId string `json:"publicip_id" xml:"publicip_id"`

	Body *CreatePublicipTagRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreatePublicipTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipTagRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicipTagRequest", string(data)}, " ")
}
