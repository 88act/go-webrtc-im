// Code generated by goctl. DO NOT EDIT.
package types

type RankResp struct {
}

type ShopResp struct {
}

type NotifyResp struct {
}

type BattleResp struct {
}

type PackResp struct {
}

type ChangeScoreReq struct {
}

type ChangeScoreResp struct {
}

type PageInfoReq struct {
	Page     int    `json:"page"`     // 页码
	PageSize int    `json:"pageSize"` //页大小
	Key      string `json:"key"`      // key
}

type OkResp struct {
	Msg   string `json:"msg"`   //  返回msg
	Value string `json:"value"` // 值
	Id    int64  `json:"id"`    // id
}

type IdReq struct {
	Id int64 `json:"id"` // id  请求
}

type ValReq struct {
	Value string `json:"value"` // 值 请求
}
