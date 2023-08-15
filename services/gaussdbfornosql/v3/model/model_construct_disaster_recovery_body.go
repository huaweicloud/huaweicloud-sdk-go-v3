package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConstructDisasterRecoveryBody struct {

	// 容灾ID。 对容灾角色为主的实例下发搭建容灾接口时不传该参数，接口成功响应后返回生成的容灾ID。 对容灾角色为备的实例下发建容灾接口时必传该参数，且必须与上述生成的容灾ID保持一致。
	Id *string `json:"id,omitempty"`

	// 搭建容灾关系的别名。
	Alias string `json:"alias"`

	// 建立容灾关系所需要的容灾密码，搭建同一容灾关系的两次调用容灾密码必须保持一致。 容灾密码为容灾集群内部数据通信所用，不能用于客户端连接使用。
	Password string `json:"password"`

	// 指定当前实例的容灾角色。取值为master或slave，表示在容灾关系中角色为主或备。
	InstanceRole string `json:"instance_role"`

	DisasterRecoveryInstance *ConstructDisasterRecoveryInstance `json:"disaster_recovery_instance"`
}

func (o ConstructDisasterRecoveryBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConstructDisasterRecoveryBody struct{}"
	}

	return strings.Join([]string{"ConstructDisasterRecoveryBody", string(data)}, " ")
}
