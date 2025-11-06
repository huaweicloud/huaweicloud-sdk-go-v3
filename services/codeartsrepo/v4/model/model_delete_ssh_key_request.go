package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSshKeyRequest Request Object
type DeleteSshKeyRequest struct {

	// **参数解释：** 资源Id，通过获取代码组权限资源点列表获取的数据中的Id
	KeyId int32 `json:"key_id"`
}

func (o DeleteSshKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSshKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSshKeyRequest", string(data)}, " ")
}
