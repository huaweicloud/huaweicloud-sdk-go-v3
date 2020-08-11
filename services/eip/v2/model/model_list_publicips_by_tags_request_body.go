/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 通过标签过滤弹性公网IP的请求体
type ListPublicipsByTagsRequestBody struct {
	// 包含标签，最多包含10个key。  每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。  Key不能重复，同一个key中values不能重复。
	Tags []TagReq `json:"tags,omitempty"`
	// 查询记录数（action为count时无此参数）
	Limit int32 `json:"limit,omitempty"`
	// 索引位置， 从offset指定的下一条数据开始查询。 查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数（action为count时无此参数）
	Offset int32 `json:"offset,omitempty"`
	// 操作标识：  filter分页查询  count查询总数
	Action string `json:"action"`
	// 搜索字段，key为要匹配的字段，当前仅支持resource_name。value为匹配的值。此字段为固定字典值。
	Matches []MatchReq `json:"matches,omitempty"`
}
