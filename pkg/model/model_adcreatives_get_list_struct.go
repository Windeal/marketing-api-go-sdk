/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type AdcreativesGetListStruct struct {
	CampaignId                    int64                            `json:"campaign_id,omitempty"`
	AdcreativeId                  int64                            `json:"adcreative_id,omitempty"`
	OuterAdcreativeId             int64                            `json:"outer_adcreative_id,omitempty"`
	AdcreativeName                string                           `json:"adcreative_name,omitempty"`
	PageType                      PageTypeRead                     `json:"page_type,omitempty"`
	PageSpec                      PageSpec                         `json:"page_spec,omitempty"`
	LinkPageType                  LinkPageType                     `json:"link_page_type,omitempty"`
	LinkNameType                  LinkNameTypeMp                   `json:"link_name_type,omitempty"`
	LinkPageSpec                  LinkPageSpec                     `json:"link_page_spec,omitempty"`
	QqMiniGameTrackingQueryString string                           `json:"qq_mini_game_tracking_query_string,omitempty"`
	DeepLinkUrl                   string                           `json:"deep_link_url,omitempty"`
	UniversalLinkUrl              string                           `json:"universal_link_url,omitempty"`
	SiteSet                       []string                         `json:"site_set,omitempty"`
	AutomaticSiteEnabled          bool                             `json:"automatic_site_enabled,omitempty"`
	PromotedObjectType            PromotedObjectType               `json:"promoted_object_type,omitempty"`
	PromotedObjectId              string                           `json:"promoted_object_id,omitempty"`
	ProfileId                     int64                            `json:"profile_id,omitempty"`
	CreatedTime                   int64                            `json:"created_time,omitempty"`
	LastModifiedTime              int64                            `json:"last_modified_time,omitempty"`
	ShareContentSpec              ShareContentSpec                 `json:"share_content_spec,omitempty"`
	DynamicAdcreativeSpec         DynamicAdcreativeSpec            `json:"dynamic_adcreative_spec,omitempty"`
	IsDeleted                     bool                             `json:"is_deleted,omitempty"`
	IsDynamicCreative             bool                             `json:"is_dynamic_creative,omitempty"`
	MultiShareOptimizationEnabled bool                             `json:"multi_share_optimization_enabled,omitempty"`
	ComponentId                   int64                            `json:"component_id,omitempty"`
	OnlineEnabled                 bool                             `json:"online_enabled,omitempty"`
	RevisedAdcreativeSpec         RevisedAdcreativeSpec            `json:"revised_adcreative_spec,omitempty"`
	Category                      []int64                          `json:"category,omitempty"`
	Label                         []string                         `json:"label,omitempty"`
	UnionMarketSwitch             bool                             `json:"union_market_switch,omitempty"`
	PlayablePageMaterialId        string                           `json:"playable_page_material_id,omitempty"`
	FeedsVideoCommentSwitch       bool                             `json:"feeds_video_comment_switch,omitempty"`
	AdcreativeTemplateId          int64                            `json:"adcreative_template_id,omitempty"`
	AdcreativeElements            AdcreativeCreativeElementsReadMp `json:"adcreative_elements,omitempty"`
}