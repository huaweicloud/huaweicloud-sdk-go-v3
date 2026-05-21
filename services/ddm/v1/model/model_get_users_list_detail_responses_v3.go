package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetUsersListDetailResponsesV3 struct {

	// **参数解释**：  DDM实例账号的名称。  **取值范围**：  不涉及。
	Name string `json:"name"`

	// **参数解释**：  DDM实例账号的状态。  **取值范围**：  不涉及。
	Status string `json:"status"`

	// **参数解释**：  DDM实例账号的基础权限。  **取值范围**：  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT
	BaseAuthority []string `json:"base_authority"`

	// **参数解释**：  DDM实例账号密码的有效期。  **取值范围**：  取值范围为0-65535的整数，单位为天。  0与空表示密码永不过期。
	PasswordLifetime *int64 `json:"password_lifetime,omitempty"`

	// **参数解释**：  DDM实例账号密码最近一次的修改时间。  格式为yyyy-mm-ddThh:mm:ssZ。其中，T指定某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。  **取值范围**：  不涉及。
	PasswordLastChanged *string `json:"password_last_changed,omitempty"`

	// **参数解释**：  账号的描述信息。  **取值范围**：  不涉及。
	Description string `json:"description"`

	// **参数解释**：  DDM实例账号的创建时间。  格式为yyyy-mm-ddThh:mm:ssZ。其中，T指定某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。  **取值范围**：  不涉及。
	Created string `json:"created"`

	// **参数解释**：  关联的逻辑库集合。账号只对已关联的逻辑库有访问权限。  **取值范围**：  不涉及。
	Databases []GetUsersListdatabaseV3 `json:"databases"`

	// **参数解释**：  DDM实例账号的创建时间。  格式为yyyy-mm-ddThh:mm:ssZ。其中，T指定某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。  **取值范围**：  不涉及。
	ExpirationTime *string `json:"expiration_time,omitempty"`
}

func (o GetUsersListDetailResponsesV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUsersListDetailResponsesV3 struct{}"
	}

	return strings.Join([]string{"GetUsersListDetailResponsesV3", string(data)}, " ")
}
