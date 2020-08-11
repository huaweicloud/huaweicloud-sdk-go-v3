/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Response Object
type ShowOnlineUsersResponse struct {
	// 查询结果的总元素数量
	Total int32 `json:"total,omitempty"`
	// 正在推流的音视频信息
	UserInfo []UserInfo `json:"user_info,omitempty"`
}
