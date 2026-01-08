package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVirsubnetRequest Request Object
type ShowVirsubnetRequest struct {

	// **参数解释**： 子网的资源ID。 **取值范围**： 不涉及。
	VirsubnetId string `json:"virsubnet_id"`
}

func (o ShowVirsubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirsubnetRequest struct{}"
	}

	return strings.Join([]string{"ShowVirsubnetRequest", string(data)}, " ")
}
