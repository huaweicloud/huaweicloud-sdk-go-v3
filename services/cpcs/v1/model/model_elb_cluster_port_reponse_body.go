package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ElbClusterPortReponseBody struct {

	// uuid
	Id *string `json:"id,omitempty"`

	// 该端口归属的集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 租户的elbId
	ElbId *string `json:"elb_id,omitempty"`

	// 租户的elb的ip
	ElbIp *string `json:"elb_ip,omitempty"`

	// PROXY：代理模式端口 TUNNEL：隧道模式自定端口 TUNNEL_FIXED：隧道模式的固定隧道端口
	Mode *ElbClusterPortReponseBodyMode `json:"mode,omitempty"`

	// elb监听端口
	ListenerPort *int32 `json:"listener_port,omitempty"`

	// elb监听器id
	ListenerId *string `json:"listener_id,omitempty"`

	// 后端服务组绑定的后端服务器的业务端口
	ServerGroupPort *int32 `json:"server_group_port,omitempty"`

	// 后端服务组id
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 最后验证时间
	ValidateTime *string `json:"validate_time,omitempty"`

	// 资源是否异常。 由于该接口功能是cpcs操作租户的elb。而租户是可以二次操作cpcs创建的端口映射的。cpcs提供了一个检测接口，用以检测cpcs创建的这一套监听端口是否有误。
	Wrong *bool `json:"wrong,omitempty"`

	// 若端口有误。即wrong=true时。该字段描述错误的地方
	WrongMsg *string `json:"wrong_msg,omitempty"`
}

func (o ElbClusterPortReponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbClusterPortReponseBody struct{}"
	}

	return strings.Join([]string{"ElbClusterPortReponseBody", string(data)}, " ")
}

type ElbClusterPortReponseBodyMode struct {
	value string
}

type ElbClusterPortReponseBodyModeEnum struct {
	TUNNEL_FIXED ElbClusterPortReponseBodyMode
	TUNNEL       ElbClusterPortReponseBodyMode
	PROXY        ElbClusterPortReponseBodyMode
}

func GetElbClusterPortReponseBodyModeEnum() ElbClusterPortReponseBodyModeEnum {
	return ElbClusterPortReponseBodyModeEnum{
		TUNNEL_FIXED: ElbClusterPortReponseBodyMode{
			value: "TUNNEL_FIXED",
		},
		TUNNEL: ElbClusterPortReponseBodyMode{
			value: "TUNNEL",
		},
		PROXY: ElbClusterPortReponseBodyMode{
			value: "PROXY",
		},
	}
}

func (c ElbClusterPortReponseBodyMode) Value() string {
	return c.value
}

func (c ElbClusterPortReponseBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ElbClusterPortReponseBodyMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
