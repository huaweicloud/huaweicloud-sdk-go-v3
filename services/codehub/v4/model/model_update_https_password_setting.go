package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpsPasswordSetting struct {

	// **参数解释：** 是否用https的认证方式 true,false。 **取值范围：** 字符串长度不少于1，不超过1000。
	HttpsCloneIamAuth *string `json:"https_clone_iam_auth,omitempty"`
}

func (o UpdateHttpsPasswordSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsPasswordSetting struct{}"
	}

	return strings.Join([]string{"UpdateHttpsPasswordSetting", string(data)}, " ")
}
