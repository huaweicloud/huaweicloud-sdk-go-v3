package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessPointResponse Response Object
type CreateAccessPointResponse struct {

	// **参数解释：** 创建的接入点ID **取值范围：** 不涉及
	AccessPointId  *string `json:"access_point_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPointResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessPointResponse", string(data)}, " ")
}
