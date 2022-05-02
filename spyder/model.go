package spyder

type AutoGenerated struct {
	Message    string `json:"message"`
	Data       Data   `json:"data"`
	StatusCode int    `json:"status_code"`
	Prompt     string `json:"prompt"`
	Time       int64  `json:"time"`
}
type Cursor struct {
	HasMore        bool `json:"has_more"`
	RefreshCursor  int  `json:"refresh_cursor"`
	LoadmoreCursor int  `json:"loadmore_cursor"`
	FeedLen        int  `json:"feed_len"`
}
type DownloadList struct {
	URL string `json:"url"`
}

type Icon struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type CategoryList struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"display_order"`
}
type BackgroundImage struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type IconImage struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type ActivityIcon struct {
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	IconImage IconImage `json:"icon_image"`
}
type BaseHashtag struct {
	Icon              Icon            `json:"icon"`
	TopGif            interface{}     `json:"top_gif"`
	CategoryList      []CategoryList  `json:"category_list"`
	ID                int             `json:"id"`
	BackgroundImage   BackgroundImage `json:"background_image"`
	ActivityIcon      ActivityIcon    `json:"activity_icon"`
	GuideWord         string          `json:"guide_word"`
	TargetTcsIDStr    interface{}     `json:"target_tcs_id_str"`
	TopStatus         int             `json:"top_status"`
	Avatars           interface{}     `json:"avatars"`
	ActionEndTime     int             `json:"action_end_time"`
	SponsorAvatar     interface{}     `json:"sponsor_avatar"`
	Tag               string          `json:"tag"`
	Extra             interface{}     `json:"extra"`
	Intro             string          `json:"intro"`
	GuideText         string          `json:"guide_text"`
	GuideLink         string          `json:"guide_link"`
	CreateTime        int             `json:"create_time"`
	Atmosphere        interface{}     `json:"atmosphere"`
	TopText           string          `json:"top_text"`
	ForbiddenPostItem bool            `json:"forbidden_post_item"`
	IDStr             string          `json:"id_str"`
	Name              string          `json:"name"`
	JumpLink          interface{}     `json:"jump_link"`
	ActionStartTime   int             `json:"action_start_time"`
	Status            int             `json:"status"`
	SponsorRankSchema interface{}     `json:"sponsor_rank_schema"`
	HashtagTopType    int             `json:"hashtag_top_type"`
	HashtagType       int             `json:"hashtag_type"`
	IntroSchema       []interface{}   `json:"intro_schema"`
	ModifyTime        int             `json:"modify_time"`
	BoostInfo         interface{}     `json:"boost_info"`
	TopItemIds        []string        `json:"top_item_ids"`
	TargetTcsID       interface{}     `json:"target_tcs_id"`
}
type HashtagSchema struct {
	Extra           interface{} `json:"extra"`
	Intro           interface{} `json:"intro"`
	IDStr           string      `json:"id_str"`
	DisplayTag      interface{} `json:"display_tag"`
	DisplayTagInfo  interface{} `json:"display_tag_info"`
	BaseHashtag     BaseHashtag `json:"base_hashtag"`
	Schema          string      `json:"schema"`
	IsFollow        bool        `json:"is_follow"`
	ShowNum         int64       `json:"show_num"`
	WorksNum        int         `json:"works_num"`
	FollowersNum    int         `json:"followers_num"`
	Tag             interface{} `json:"tag"`
	EnterNum        int         `json:"enter_num"`
	Icon            interface{} `json:"icon"`
	Status          interface{} `json:"status"`
	BackgroundImage interface{} `json:"background_image"`
}
type CoverImage struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type URLList struct {
	URL     string `json:"url"`
	Expires int    `json:"expires"`
}
type VideoDownload struct {
	AnimatedCoverImage interface{} `json:"animated_cover_image"`
	URI                string      `json:"uri"`
	Width              int         `json:"width"`
	CoverImage         CoverImage  `json:"cover_image"`
	URLList            []URLList   `json:"url_list"`
	AlarmText          interface{} `json:"alarm_text"`
	Height             int         `json:"height"`
	CodecType          int         `json:"codec_type"`
	Duration           float64     `json:"duration"`
	Definition         int         `json:"definition"`
	VideoModel         string      `json:"video_model"`
	P2PType            interface{} `json:"p2p_type"`
	FileHash           interface{} `json:"file_hash"`
}
type VideoHigh struct {
	AnimatedCoverImage interface{} `json:"animated_cover_image"`
	URI                string      `json:"uri"`
	Width              int         `json:"width"`
	CoverImage         CoverImage  `json:"cover_image"`
	URLList            []URLList   `json:"url_list"`
	AlarmText          interface{} `json:"alarm_text"`
	Height             int         `json:"height"`
	CodecType          int         `json:"codec_type"`
	Duration           float64     `json:"duration"`
	Definition         int         `json:"definition"`
	VideoModel         string      `json:"video_model"`
	P2PType            int         `json:"p2p_type"`
	FileHash           interface{} `json:"file_hash"`
}
type VideoFallback struct {
	AnimatedCoverImage interface{}   `json:"animated_cover_image"`
	URI                string        `json:"uri"`
	Width              int           `json:"width"`
	CoverImage         CoverImage    `json:"cover_image"`
	URLList            []interface{} `json:"url_list"`
	AlarmText          interface{}   `json:"alarm_text"`
	Height             int           `json:"height"`
	CodecType          interface{}   `json:"codec_type"`
	Duration           float64       `json:"duration"`
	Definition         interface{}   `json:"definition"`
	VideoModel         string        `json:"video_model"`
	P2PType            interface{}   `json:"p2p_type"`
	FileHash           interface{}   `json:"file_hash"`
}
type Video struct {
	VideoID             string          `json:"video_id"`
	VideoHeight         int             `json:"video_height"`
	HashtagSchema       []HashtagSchema `json:"hashtag_schema"`
	VideoDownload       VideoDownload   `json:"video_download"`
	VideoMid            interface{}     `json:"video_mid"`
	CoverImage          CoverImage      `json:"cover_image"`
	Transport           bool            `json:"transport"`
	Title               interface{}     `json:"title"`
	VideoLow            interface{}     `json:"video_low"`
	VideoHigh           VideoHigh       `json:"video_high"`
	VideoFallback       VideoFallback   `json:"video_fallback"`
	Text                string          `json:"text"`
	VideoGodCommentUrls interface{}     `json:"video_god_comment_urls"`
	TransportText       string          `json:"transport_text"`
	Duration            float64         `json:"duration"`
	VideoWidth          int             `json:"video_width"`
	SupportLivePhoto    bool            `json:"support_live_photo"`
	Animate             interface{}     `json:"animate"`
	TailAdPassthrough   string          `json:"tail_ad_passthrough"`
}

type Cover struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type CellUICtrl struct {
	DisplayStyle   interface{} `json:"display_style"`
	ShowAuthorTags bool        `json:"show_author_tags"`
	ImageStyle     interface{} `json:"image_style"`
}
type Avatar struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type WardUsers struct {
	CommercePermissionList interface{} `json:"commerce_permission_list"`
	Region                 string      `json:"region"`
	ID                     int64       `json:"id"`
	Level                  int         `json:"level"`
	DecorationList         interface{} `json:"decoration_list"`
	IsFollowed             interface{} `json:"is_followed"`
	RecommendReason        interface{} `json:"recommend_reason"`
	CertifyInfo            interface{} `json:"certify_info"`
	AuthorInfo             interface{} `json:"author_info"`
	Horoscope              interface{} `json:"horoscope"`
	FollowersCount         interface{} `json:"followers_count"`
	Description            string      `json:"description"`
	HideAge                interface{} `json:"hide_age"`
	LikeCount              interface{} `json:"like_count"`
	Punishments            interface{} `json:"punishments"`
	IDStr                  string      `json:"id_str"`
	Name                   string      `json:"name"`
	Achievements           interface{} `json:"achievements"`
	Status                 interface{} `json:"status"`
	Avatar                 Avatar      `json:"avatar"`
	CreativeLevelInfo      interface{} `json:"creative_level_info"`
	IsFollowing            interface{} `json:"is_following"`
	ProfileSchema          string      `json:"profile_schema"`
	FollowingsCount        interface{} `json:"followings_count"`
	Gender                 interface{} `json:"gender"`
	Age                    interface{} `json:"age"`
	LiveAuth               bool        `json:"live_auth"`
	BroadcastInfo          interface{} `json:"broadcast_info"`
	RecommendTag           interface{} `json:"recommend_tag"`
	GodCommentCount        interface{} `json:"god_comment_count"`
	Language               string      `json:"language"`
	InteractionLimitation  interface{} `json:"interaction_limitation"`
	VoteCount              int         `json:"vote_count"`
}
type WardInfo struct {
	IsWard        bool          `json:"is_ward"`
	IsVisible     bool          `json:"is_visible"`
	WardReplyID   int           `json:"ward_reply_id"`
	WardCount     int           `json:"ward_count"`
	TrendMessages []interface{} `json:"trend_messages"`
	WardCommentID int           `json:"ward_comment_id"`
	WardUsers     []WardUsers   `json:"ward_users"`
	HasNew        bool          `json:"has_new"`
}
type Privilege struct {
	BulletPlay bool `json:"bullet_play"`
	BulletPost bool `json:"bullet_post"`
}
type Ios struct {
	QqStrategy      int `json:"qq_strategy"`
	WeixinStrategy  int `json:"weixin_strategy"`
	Channel         int `json:"channel"`
	QzoneStrategy   int `json:"qzone_strategy"`
	MomentsStrategy int `json:"moments_strategy"`
}
type Android struct {
	QqStrategy      int `json:"qq_strategy"`
	WeixinStrategy  int `json:"weixin_strategy"`
	Channel         int `json:"channel"`
	QzoneStrategy   int `json:"qzone_strategy"`
	MomentsStrategy int `json:"moments_strategy"`
}
type Share struct {
	ShareText       string  `json:"share_text"`
	MomentsStrategy int     `json:"moments_strategy"`
	QzoneURL        string  `json:"qzone_url"`
	QqStrategy      int     `json:"qq_strategy"`
	Ios             Ios     `json:"ios"`
	ImageURL        string  `json:"image_url"`
	CompoundPageURL string  `json:"compound_page_url"`
	Schema          string  `json:"schema"`
	LinkText        string  `json:"link_text"`
	Title           string  `json:"title"`
	Android         Android `json:"android"`
	WeixinStrategy  int     `json:"weixin_strategy"`
	WechatURL       string  `json:"wechat_url"`
	QqURL           string  `json:"qq_url"`
	MomentsURL      string  `json:"moments_url"`
	ShareURL        string  `json:"share_url"`
	LargeImageURL   string  `json:"large_image_url"`
	QzoneStrategy   int     `json:"qzone_strategy"`
	Content         string  `json:"content"`
}
type DecorationInfos struct {
	DecorationType int    `json:"decoration_type"`
	Icon           string `json:"icon"`
	Description    string `json:"description"`
	Schema         string `json:"schema"`
	DecorationID   int    `json:"decoration_id"`
}
type DecorationList struct {
	Category        int               `json:"category"`
	DecorationInfos []DecorationInfos `json:"decoration_infos"`
}
type Author struct {
	CommercePermissionList interface{}      `json:"commerce_permission_list"`
	Region                 string           `json:"region"`
	ID                     int64            `json:"id"`
	Level                  int              `json:"level"`
	DecorationList         []DecorationList `json:"decoration_list"`
	IsFollowed             interface{}      `json:"is_followed"`
	RecommendReason        interface{}      `json:"recommend_reason"`
	CertifyInfo            interface{}      `json:"certify_info"`
	AuthorInfo             interface{}      `json:"author_info"`
	Horoscope              string           `json:"horoscope"`
	FollowersCount         int              `json:"followers_count"`
	Description            string           `json:"description"`
	HideAge                interface{}      `json:"hide_age"`
	LikeCount              int              `json:"like_count"`
	Punishments            interface{}      `json:"punishments"`
	IDStr                  string           `json:"id_str"`
	Name                   string           `json:"name"`
	Achievements           []interface{}    `json:"achievements"`
	Status                 int              `json:"status"`
	Avatar                 Avatar           `json:"avatar"`
	CreativeLevelInfo      interface{}      `json:"creative_level_info"`
	IsFollowing            bool             `json:"is_following"`
	ProfileSchema          string           `json:"profile_schema"`
	FollowingsCount        int              `json:"followings_count"`
	Gender                 int              `json:"gender"`
	Age                    string           `json:"age"`
	LiveAuth               bool             `json:"live_auth"`
	BroadcastInfo          interface{}      `json:"broadcast_info"`
	RecommendTag           interface{}      `json:"recommend_tag"`
	GodCommentCount        int              `json:"god_comment_count"`
	Language               string           `json:"language"`
	InteractionLimitation  interface{}      `json:"interaction_limitation"`
	VoteCount              int              `json:"vote_count"`
}
type InteractEgg struct {
	LikeStyle string `json:"like_style"`
}
type ItemRelation struct {
	IsBury     bool `json:"is_bury"`
	DiggType   int  `json:"digg_type"`
	BuryType   int  `json:"bury_type"`
	IsLike     bool `json:"is_like"`
	IsFavorite bool `json:"is_favorite"`
}
type BuryCounts struct {
	BuryCount int `json:"bury_count"`
	BuryType  int `json:"bury_type"`
}
type Stats struct {
	CommentCount    int           `json:"comment_count"`
	DubbingCount    int           `json:"dubbing_count"`
	BuryCount       int           `json:"bury_count"`
	GoDetailCount   int           `json:"go_detail_count"`
	BuryCounts      []BuryCounts  `json:"bury_counts"`
	ImpressionCount int           `json:"impression_count"`
	BulletCount     int           `json:"bullet_count"`
	ViewCount       int           `json:"view_count"`
	LikeCount       int           `json:"like_count"`
	DiggCounts      []interface{} `json:"digg_counts"`
	ShareCount      int           `json:"share_count"`
	PlayCount       int           `json:"play_count"`
}

type OriginVideoDownload struct {
	AnimatedCoverImage interface{} `json:"animated_cover_image"`
	URI                string      `json:"uri"`
	Width              int         `json:"width"`
	CoverImage         CoverImage  `json:"cover_image"`
	URLList            []URLList   `json:"url_list"`
	AlarmText          interface{} `json:"alarm_text"`
	Height             int         `json:"height"`
	CodecType          int         `json:"codec_type"`
	Duration           float64     `json:"duration"`
	Definition         int         `json:"definition"`
	VideoModel         string      `json:"video_model"`
	P2PType            interface{} `json:"p2p_type"`
	FileHash           interface{} `json:"file_hash"`
}

type AhaImage struct {
	IsGif        bool           `json:"is_gif"`
	DownloadList []DownloadList `json:"download_list"`
	URLList      []URLList      `json:"url_list"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	URI          string         `json:"uri"`
}
type Item struct {
	Stage                interface{}         `json:"stage"`
	Video                Video               `json:"video"`
	ClubInfo             interface{}         `json:"club_info"`
	Activity             interface{}         `json:"activity"`
	HumanTags            interface{}         `json:"human_tags"`
	DupItemSchema        interface{}         `json:"dup_item_schema"`
	ShowFeatureDigg      bool                `json:"show_feature_digg"`
	CommentPostItem      bool                `json:"comment_post_item"`
	Cover                Cover               `json:"cover"`
	AiTag                interface{}         `json:"ai_tag"`
	RecommendTags        interface{}         `json:"recommend_tags"`
	IsOrigin             interface{}         `json:"is_origin"`
	ItemIDStr            string              `json:"item_id_str"`
	Link                 interface{}         `json:"link"`
	ItemCellType         int                 `json:"item_cell_type"`
	InteractionStatus    interface{}         `json:"interaction_status"`
	AuthorTags           string              `json:"author_tags"`
	QualityLevel         interface{}         `json:"quality_level"`
	Position             interface{}         `json:"position"`
	RelatedCommentID     int                 `json:"related_comment_id"`
	AncestorCommentID    interface{}         `json:"ancestor_comment_id"`
	Note                 interface{}         `json:"note"`
	HotInfo              interface{}         `json:"hot_info"`
	EpisodeInfo          interface{}         `json:"episode_info"`
	Rating               int                 `json:"rating"`
	CellUICtrl           CellUICtrl          `json:"cell_ui_ctrl"`
	VideoType            int                 `json:"video_type"`
	DebugInfo            interface{}         `json:"debug_info"`
	Status               int                 `json:"status"`
	AncestorSchema       interface{}         `json:"ancestor_schema"`
	WardInfo             WardInfo            `json:"ward_info"`
	RelatedID            interface{}         `json:"related_id"`
	RecreateMetaInfoList []interface{}       `json:"recreate_meta_info_list"`
	Source               interface{}         `json:"source"`
	Privilege            Privilege           `json:"privilege"`
	JumpLink             interface{}         `json:"jump_link"`
	Share                Share               `json:"share"`
	MappingGids          interface{}         `json:"mapping_gids"`
	Animations           interface{}         `json:"animations"`
	RelatedCommentIDStr  string              `json:"related_comment_id_str"`
	MicroAppID           interface{}         `json:"micro_app_id"`
	ReviewStatus         interface{}         `json:"review_status"`
	CreateTime           int                 `json:"create_time"`
	ShowFeatureBury      bool                `json:"show_feature_bury"`
	OriginVideoID        string              `json:"origin_video_id"`
	AdminDebug           interface{}         `json:"admin_debug"`
	UserStatus           interface{}         `json:"user_status"`
	GameCardInfo         interface{}         `json:"game_card_info"`
	Author               Author              `json:"author"`
	PromotionInfo        interface{}         `json:"promotion_info"`
	ClubRelation         interface{}         `json:"club_relation"`
	InteractEgg          InteractEgg         `json:"interact_egg"`
	LogItemExtra         string              `json:"log_item_extra"`
	AuthorDisplayTags    interface{}         `json:"author_display_tags"`
	Comments             []interface{}       `json:"comments"`
	TextSchema           interface{}         `json:"text_schema"`
	Duration             float64             `json:"duration"`
	AuditInfo            interface{}         `json:"audit_info"`
	ItemRelation         ItemRelation        `json:"item_relation"`
	Stats                Stats               `json:"stats"`
	Content              string              `json:"content"`
	DrainageInfo         interface{}         `json:"drainage_info"`
	CanDownload          bool                `json:"can_download"`
	Level                interface{}         `json:"level"`
	PostSource           int                 `json:"post_source"`
	AlbumIntro           interface{}         `json:"album_intro"`
	NewAppVisible        interface{}         `json:"new_app_visible"`
	LastCommentTime      interface{}         `json:"last_comment_time"`
	FrozenToast          string              `json:"frozen_toast"`
	AncestorID           string              `json:"ancestor_id"`
	ExpireTime           interface{}         `json:"expire_time"`
	VitalComments        interface{}         `json:"vital_comments"`
	OriginVideoDownload  OriginVideoDownload `json:"origin_video_download"`
	AhaImage             []AhaImage          `json:"aha_image"`
	DefaultSchema        string              `json:"default_schema"`
	NeihanStyle          int                 `json:"neihan_style"`
	ItemID               int64               `json:"item_id"`
	ItemType             int                 `json:"item_type"`
}
type DislikeOptions struct {
	OptionID    int       `json:"option_id"`
	DislikeType int       `json:"dislike_type"`
	IconImage   IconImage `json:"icon_image"`
	DislikeDesc string    `json:"dislike_desc"`
}
type BoardInfo struct {
	GodRank   int `json:"god_rank"`
	BlockType int `json:"block_type"`
}
type Data struct {
	FollowOptions    interface{}      `json:"follow_options"`
	CollectionInfo   interface{}      `json:"collection_info"`
	BlockInfo        interface{}      `json:"block_info"`
	PassThrough      string           `json:"pass_through"`
	Item             Item             `json:"item"`
	Stickup          bool             `json:"stickup"`
	LastViewTime     int              `json:"last_view_time"`
	UserItems        interface{}      `json:"user_items"`
	SearchAttachInfo interface{}      `json:"search_attach_info"`
	LiveInfo         interface{}      `json:"live_info"`
	CellType         int              `json:"cell_type"`
	Hashtag          interface{}      `json:"hashtag"`
	BannerInfo       interface{}      `json:"banner_info"`
	CellIDStr        string           `json:"cell_id_str"`
	Button           interface{}      `json:"button"`
	ReplyInfo        interface{}      `json:"reply_info"`
	DislikeOptions   []DislikeOptions `json:"dislike_options"`
	CommentShowAd    bool             `json:"comment_show_ad"`
	AlbumInfo        interface{}      `json:"album_info"`
	CommentInfo      interface{}      `json:"comment_info"`
	CellID           int64            `json:"cell_id"`
	DisplayTime      int              `json:"display_time"`
	UserInfo         interface{}      `json:"user_info"`
	UserHotItemList  interface{}      `json:"user_hot_item_list"`
	UsersHashtag     interface{}      `json:"users_hashtag"`
	ClubInfo         interface{}      `json:"club_info"`
	AdInfo           interface{}      `json:"ad_info"`
	SaasLiveInfo     interface{}      `json:"saas_live_info"`
	Comment          interface{}      `json:"comment"`
	BoardInfo        BoardInfo        `json:"board_info"`
}
