package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirsubnetCidrReservationRequestBody This is a auto create Body Object
type CreateVirsubnetCidrReservationRequestBody struct {

	// **参数解释**： 是否只预检此次请求。 **约束限制**： 不涉及。 **取值范围**： - true：发送检查请求，不会创建子网预留网段。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。 - false：发送正常请求，并直接创建子网预留网段。 **默认取值**： false
	DryRun *bool `json:"dry_run,omitempty"`

	VirsubnetCidrReservation *CreateVirsubnetCidrReservationOption `json:"virsubnet_cidr_reservation"`
}

func (o CreateVirsubnetCidrReservationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirsubnetCidrReservationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVirsubnetCidrReservationRequestBody", string(data)}, " ")
}
