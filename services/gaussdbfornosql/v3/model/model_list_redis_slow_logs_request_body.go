package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRedisSlowLogsRequestBody struct {

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：开始时间不得早于当前时间30天。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。注：结束时间不能晚于当前时间。
	EndTime string `json:"end_time"`

	// 表示每次查询的日志条数，最大限制100条。
	Limit int32 `json:"limit"`

	// 日志单行序列号，第一次查询时不需要此参数，下一次查询时需要使用，可从上一次查询的返回信息中获取。 说明：当次查询从line_num的下一条日志开始查询，不包含当前line_num日志。
	LineNum *string `json:"line_num,omitempty"`

	// 语句类型，取空值，表示查询所有语句类型。支持查询的所有语句类型如下（以\"|\"分割）： set|get|del|incr|incrby|incrbyfloat|decr|decrby|getset|append|mget|keys|setnx|setex|psetex|delvx|mset| msetnx|getrange|substr|setrange|strlen|exists|expire|pexpire|expireat|pexpireat|ttl|pttl|persist|type| scanx|pksetexat|sort|hdel|hset|hget|hgetall|hexists|hincrby|hincrbyfloat|hkeys|hlen|hmget|hmset|hsetnx| hstrlen|hvals|hscan|hscanx|pkhscanrange|pkhrscanrange|lindex|linsert|llen|lpop|lpush|lpushx|lrange|lrem| lset|ltrim|rpop|rpoplpush|rpush|rpushx|zadd|zcard|zscan|zincrby|zrange|zrevrange|zrangebyscore| zrevrangebyscore|zcount|zrem|zunionstore|zinterstore|zrank|zrevrank|zscore|zrangebylex|zrevrangebylex| zlexcount|zremrangebyrank|zremrangebyscore|zremrangebylex|zpopmax|zpopmin|sadd|spop|scard|smembers|sscan| srem|sunion|sunionstore|sinter|sinterstore|sismember|sdiff|sdiffstore|smove|srandmember|bitset|bitget| bitcount|bitpos|bitop|bitfield|pfadd|pfcount|pfmerge|geoadd|georadiusbymember|georadius|geohash|geodist| geopos|xadd|xack|xgroup|xdel|xtrim|xlen|xrange|xrevrange|xclaim|xpending|xinfo|xread|xreadgroup|
	OperateType *string `json:"operate_type,omitempty"`

	// 节点ID，取空值，表示查询实例下所有允许查询的节点。 具体取值请参考查询实例列表和详情接口\"ListInstances\"中nodes字段数据结构说明的“id”。
	NodeId *string `json:"node_id,omitempty"`

	// 根据多个关键字搜索日志全文，表示同时匹配所有关键字。 - 最多支持10个关键字。 - 每个关键字最大长度不超过512个字符。
	Keywords *[]string `json:"keywords,omitempty"`

	// 支持根据最大执行时间范围查找日志。单位：ms
	MaxCostTime *float64 `json:"max_cost_time,omitempty"`

	// 支持根据最小执行时间范围查找日志。单位：ms
	MinCostTime *float64 `json:"min_cost_time,omitempty"`
}

func (o ListRedisSlowLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedisSlowLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListRedisSlowLogsRequestBody", string(data)}, " ")
}
