package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ClientInfo struct {

	// 客户端ID
	Id *string `json:"id,omitempty"`

	// 客户端的地址和端口
	Addr *string `json:"addr,omitempty"`

	// 套接字所使用的文件描述符。
	Fd *string `json:"fd,omitempty"`

	// 客户端的名称
	Name *string `json:"name,omitempty"`

	// 最近一次执行的命令
	Cmd *string `json:"cmd,omitempty"`

	// 已连接时长（单位：秒）
	Age *int32 `json:"age,omitempty"`

	// 空闲时长（单位：秒）
	Idle *int32 `json:"idle,omitempty"`

	// 该客户端正在使用的数据库 ID
	Db *string `json:"db,omitempty"`

	// 客户端标志
	Flags *string `json:"flags,omitempty"`

	// 已订阅频道的数量
	Sub *int32 `json:"sub,omitempty"`

	// 已订阅模式的数量
	Psub *int32 `json:"psub,omitempty"`

	// 在事务中被执行的命令数量
	Multi *int32 `json:"multi,omitempty"`

	// 查询缓冲区的长度（单位为字节， 0 表示没有分配查询缓冲区）
	Qbuf *int32 `json:"qbuf,omitempty"`

	// 查询缓冲区剩余空间的长度（单位为字节， 0 表示没有剩余空间）
	QbufFree *int32 `json:"qbuf_free,omitempty"`

	// 输出缓冲区的长度（单位为字节， 0 表示没有分配输出缓冲区）
	Obl *int32 `json:"obl,omitempty"`

	// 输出列表包含的对象数量（当输出缓冲区没有剩余空间时，命令回复会以字符串对象的形式被入队到这个队列里）
	Oll *int32 `json:"oll,omitempty"`

	// 输出缓冲区和输出列表占用的内存总量
	Omem *int32 `json:"omem,omitempty"`

	// 文件描述符事件
	Events *ClientInfoEvents `json:"events,omitempty"`

	// 客户端所使用的网络类型。
	Network *string `json:"network,omitempty"`

	// 单机，主备和cluster实例地址和端口。
	Peer *string `json:"peer,omitempty"`

	// 客户端用户。
	User *string `json:"user,omitempty"`
}

func (o ClientInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientInfo struct{}"
	}

	return strings.Join([]string{"ClientInfo", string(data)}, " ")
}

type ClientInfoEvents struct {
	value string
}

type ClientInfoEventsEnum struct {
	R ClientInfoEvents
	W ClientInfoEvents
}

func GetClientInfoEventsEnum() ClientInfoEventsEnum {
	return ClientInfoEventsEnum{
		R: ClientInfoEvents{
			value: "r",
		},
		W: ClientInfoEvents{
			value: "w",
		},
	}
}

func (c ClientInfoEvents) Value() string {
	return c.value
}

func (c ClientInfoEvents) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClientInfoEvents) UnmarshalJSON(b []byte) error {
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
