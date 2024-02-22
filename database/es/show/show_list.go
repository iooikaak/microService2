package show

import "github.com/iooikaak/frame/util"

const (
	ShowListElasticSearchIndex = "show_list_v5"
	ShowListElasticSearchType  = "_doc"
)

type ShowListEsModel struct {
	StdShowId            string        `json:"std_show_id"`
	BizShowId            string        `json:"biz_show_id"`
	BizArtistIds         string        `json:"biz_artist_ids"`
	IdentityRequiredType int64         `json:"identity_required_type"`
	IsShowStdContent     bool          `json:"is_show_std_content"`
	LastShowTime         util.JsonTime `json:"last_show_time"`
}

type ShowListSearchEsModel struct {
	StdShowId            string `json:"std_show_id"`
	BizShowId            string `json:"biz_show_id"`
	BizArtistIds         string `json:"biz_artist_ids"`
	IdentityRequiredType int64  `json:"identity_required_type"`
	IsShowStdContent     bool   `json:"is_show_std_content"`
	LastShowTime         string `json:"last_show_time"`
	SortBy               string
	SortDirection        string
	From                 int
	PageSize             int
}
