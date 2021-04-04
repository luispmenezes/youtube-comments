package external

type CommentResponse struct {
	XSRFToken string      `json:"xsrf_token"`
	Endpoint  interface{} `json:"endpoint"`
	Response  Response    `json:"response"`
	Timing    Timing      `json:"timing"`
}
type Params struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type ServiceTrackingParams struct {
	Service string   `json:"service"`
	Params  []Params `json:"params"`
}
type MainAppWebResponseContext struct {
	LoggedOut bool `json:"loggedOut"`
}
type YtConfigData struct {
	VisitorData string `json:"visitorData"`
}
type WebResponseContextExtensionData struct {
	YtConfigData YtConfigData `json:"ytConfigData"`
	HasDecorated bool         `json:"hasDecorated"`
}
type ResponseContext struct {
	ServiceTrackingParams           []ServiceTrackingParams         `json:"serviceTrackingParams"`
	MainAppWebResponseContext       MainAppWebResponseContext       `json:"mainAppWebResponseContext"`
	WebResponseContextExtensionData WebResponseContextExtensionData `json:"webResponseContextExtensionData"`
}
type AuthorText struct {
	SimpleText string `json:"simpleText"`
}
type Thumbnails struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData,omitempty"`
	Label             string            `json:"label,omitempty"`
}
type AuthorThumbnail struct {
	Thumbnails    []Thumbnails  `json:"thumbnails"`
	Accessibility Accessibility `json:"accessibility"`
}
type WebCommandMetadata struct {
	URL         string `json:"url"`
	WebPageType string `json:"webPageType"`
	RootVe      int    `json:"rootVe"`
	ApiUrl      string `json:"apiUrl"`
}
type CommandMetadata struct {
	WebCommandMetadata WebCommandMetadata `json:"webCommandMetadata"`
}
type BrowseEndpoint struct {
	BrowseId         string `json:"browseId"`
	CanonicalBaseUrl string `json:"canonicalBaseUrl"`
}
type AuthorEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	BrowseEndpoint      BrowseEndpoint  `json:"browseEndpoint"`
}
type ContentText struct {
	Runs []Runs `json:"runs"`
}
type WatchEndpoint struct {
	VideoId string `json:"videoId"`
	Params  string `json:"params,omitempty"`
}
type NavigationEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	WatchEndpoint       WatchEndpoint   `json:"watchEndpoint,omitempty"`
	SignInEndpoint      SignInEndpoint  `json:"signInEndpoint,omitempty"`
}
type Runs struct {
	Text               string             `json:"text"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint,omitempty"`
}
type PublishedTimeText struct {
	Runs []Runs `json:"runs"`
}
type Style struct {
	StyleType string `json:"styleType"`
}
type Size struct {
	SizeType string `json:"sizeType"`
}
type DefaultIcon struct {
	IconType string `json:"iconType"`
}
type ToggledStyle struct {
	StyleType string `json:"styleType"`
}
type NextEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	WatchEndpoint       WatchEndpoint   `json:"watchEndpoint"`
}
type SignInEndpoint struct {
	NextEndpoint NextEndpoint `json:"nextEndpoint"`
}
type DefaultNavigationEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	SignInEndpoint      SignInEndpoint  `json:"signInEndpoint"`
}
type AccessibilityData struct {
	InnerAccessibilityData struct {
		Label string `json:"label,omitempty"`
	} `json:"accessibilityData,omitempty"`
	Label string `json:"label,omitempty"`
}
type ToggledAccessibilityData struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}
type ToggleButtonRenderer struct {
	Style                     Style                     `json:"style"`
	Size                      Size                      `json:"size"`
	IsToggled                 bool                      `json:"isToggled"`
	IsDisabled                bool                      `json:"isDisabled"`
	DefaultIcon               DefaultIcon               `json:"defaultIcon"`
	TrackingParams            string                    `json:"trackingParams"`
	DefaultTooltip            string                    `json:"defaultTooltip"`
	ToggledTooltip            string                    `json:"toggledTooltip"`
	ToggledStyle              ToggledStyle              `json:"toggledStyle"`
	DefaultNavigationEndpoint DefaultNavigationEndpoint `json:"defaultNavigationEndpoint"`
	AccessibilityData         AccessibilityData         `json:"accessibilityData"`
	ToggledAccessibilityData  ToggledAccessibilityData  `json:"toggledAccessibilityData"`
}
type LikeButton struct {
	ToggleButtonRenderer ToggleButtonRenderer `json:"toggleButtonRenderer"`
}
type Text struct {
	Runs []Runs `json:"runs"`
}
type ButtonRenderer struct {
	Style              string             `json:"style,omitempty"`
	Size               string             `json:"size,omitempty"`
	Text               Text               `json:"text"`
	Icon               Icon               `json:"icon,omitempty"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint,omitempty"`
	Accessibility      Accessibility      `json:"accessibility,omitempty"`
	TrackingParams     string             `json:"trackingParams"`
	IconPosition       string             `json:"iconPosition,omitempty"`
}
type ReplyButton struct {
	ButtonRenderer ButtonRenderer `json:"buttonRenderer"`
}
type DislikeButton struct {
	ToggleButtonRenderer ToggleButtonRenderer `json:"toggleButtonRenderer"`
}
type CreatorThumbnail struct {
	Thumbnails    []Thumbnails  `json:"thumbnails"`
	Accessibility Accessibility `json:"accessibility"`
}
type HeartIcon struct {
	IconType string `json:"iconType"`
}
type BasicColorPaletteData struct {
	ForegroundTitleColor int64 `json:"foregroundTitleColor"`
}
type HeartColor struct {
	BasicColorPaletteData BasicColorPaletteData `json:"basicColorPaletteData"`
}
type HeartedAccessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}
type CreatorHeartRenderer struct {
	CreatorThumbnail        CreatorThumbnail     `json:"creatorThumbnail"`
	HeartIcon               HeartIcon            `json:"heartIcon"`
	HeartColor              HeartColor           `json:"heartColor"`
	HeartedTooltip          string               `json:"heartedTooltip"`
	IsHearted               bool                 `json:"isHearted"`
	IsEnabled               bool                 `json:"isEnabled"`
	HeartedAccessibility    HeartedAccessibility `json:"heartedAccessibility"`
	KennedyHeartColorString string               `json:"kennedyHeartColorString"`
}
type CreatorHeart struct {
	CreatorHeartRenderer CreatorHeartRenderer `json:"creatorHeartRenderer"`
}
type CommentActionButtonsRenderer struct {
	LikeButton      LikeButton    `json:"likeButton"`
	ReplyButton     ReplyButton   `json:"replyButton"`
	DislikeButton   DislikeButton `json:"dislikeButton"`
	TrackingParams  string        `json:"trackingParams"`
	CreatorHeart    CreatorHeart  `json:"creatorHeart"`
	ProtoCreationMs string        `json:"protoCreationMs"`
	Style           string        `json:"style"`
}
type ActionButtons struct {
	CommentActionButtonsRenderer CommentActionButtonsRenderer `json:"commentActionButtonsRenderer"`
}
type VoteCount struct {
	Accessibility Accessibility `json:"accessibility"`
	SimpleText    string        `json:"simpleText"`
}
type Icon struct {
	IconType string `json:"iconType"`
}
type Label struct {
	Runs []Runs `json:"runs"`
}
type Color struct {
	BasicColorPaletteData BasicColorPaletteData `json:"basicColorPaletteData"`
}
type PinnedCommentBadgeRenderer struct {
	Icon  Icon  `json:"icon"`
	Label Label `json:"label"`
	Color Color `json:"color"`
}
type PinnedCommentBadge struct {
	PinnedCommentBadgeRenderer PinnedCommentBadgeRenderer `json:"pinnedCommentBadgeRenderer"`
}
type ExpandButton struct {
	ButtonRenderer ButtonRenderer `json:"buttonRenderer"`
}
type CollapseButton struct {
	ButtonRenderer ButtonRenderer `json:"buttonRenderer"`
}
type Visibility struct {
	Types string `json:"types"`
}
type LoggingDirectives struct {
	TrackingParams string     `json:"trackingParams"`
	Visibility     Visibility `json:"visibility"`
}
type CommentRenderer struct {
	AuthorText           AuthorText         `json:"authorText"`
	AuthorThumbnail      AuthorThumbnail    `json:"authorThumbnail"`
	AuthorEndpoint       AuthorEndpoint     `json:"authorEndpoint"`
	ContentText          ContentText        `json:"contentText"`
	PublishedTimeText    PublishedTimeText  `json:"publishedTimeText"`
	IsLiked              bool               `json:"isLiked"`
	LikeCount            int                `json:"likeCount"`
	CommentId            string             `json:"commentId"`
	ActionButtons        ActionButtons      `json:"actionButtons"`
	AuthorIsChannelOwner bool               `json:"authorIsChannelOwner"`
	VoteStatus           string             `json:"voteStatus"`
	TrackingParams       string             `json:"trackingParams"`
	VoteCount            VoteCount          `json:"voteCount"`
	PinnedCommentBadge   PinnedCommentBadge `json:"pinnedCommentBadge"`
	ExpandButton         ExpandButton       `json:"expandButton"`
	CollapseButton       CollapseButton     `json:"collapseButton"`
	ReplyCount           int                `json:"replyCount"`
	LoggingDirectives    LoggingDirectives  `json:"loggingDirectives"`
}
type Comment struct {
	CommentRenderer CommentRenderer `json:"commentRenderer"`
}
type NextContinuationData struct {
	Continuation        string `json:"continuation"`
	ClickTrackingParams string `json:"clickTrackingParams"`
	Label               Label  `json:"label"`
}
type Continuations struct {
	NextContinuationData NextContinuationData `json:"nextContinuationData"`
}
type ViewReplies struct {
	ButtonRenderer ButtonRenderer `json:"buttonRenderer"`
}
type HideReplies struct {
	ButtonRenderer ButtonRenderer `json:"buttonRenderer"`
}
type CommentRepliesRenderer struct {
	Continuations  []Continuations `json:"continuations"`
	TrackingParams string          `json:"trackingParams"`
	ViewReplies    ViewReplies     `json:"viewReplies"`
	HideReplies    HideReplies     `json:"hideReplies"`
}
type Replies struct {
	CommentRepliesRenderer CommentRepliesRenderer `json:"commentRepliesRenderer"`
}
type CommentThreadRenderer struct {
	Comment               Comment           `json:"comment"`
	Replies               Replies           `json:"replies,omitempty"`
	TrackingParams        string            `json:"trackingParams"`
	RenderingPriority     string            `json:"renderingPriority"`
	IsModeratedElqComment bool              `json:"isModeratedElqComment"`
	LoggingDirectives     LoggingDirectives `json:"loggingDirectives"`
}
type CountText struct {
	Runs []Runs `json:"runs"`
}
type PlaceHolderText struct {
	Runs []Runs `json:"runs"`
}
type PrepareAccountEndpoint struct {
	ClickTrackingParams string          `json:"clickTrackingParams"`
	CommandMetadata     CommandMetadata `json:"commandMetadata"`
	SignInEndpoint      SignInEndpoint  `json:"signInEndpoint"`
}
type CommentSimpleBoxRenderer struct {
	AuthorThumbnail        AuthorThumbnail        `json:"authorThumbnail"`
	PlaceHolderText        PlaceHolderText        `json:"placeholderText"`
	PrepareAccountEndpoint PrepareAccountEndpoint `json:"prepareAccountEndpoint"`
	TrackingParams         string                 `json:"trackingParams"`
	AvatarSize             string                 `json:"avatarSize"`
}
type CreateRenderer struct {
	CommentSimpleBoxRenderer CommentSimpleBoxRenderer `json:"commentSimpleboxRenderer"`
}
type ReloadContinuationData struct {
	Continuation        string `json:"continuation"`
	ClickTrackingParams string `json:"clickTrackingParams"`
}
type Continuation struct {
	ReloadContinuationData ReloadContinuationData `json:"reloadContinuationData"`
}
type SubMenuItems struct {
	Title          string       `json:"title"`
	Selected       bool         `json:"selected"`
	Continuation   Continuation `json:"continuation"`
	TrackingParams string       `json:"trackingParams"`
}
type SortFilterSubMenuRenderer struct {
	SubMenuItems   []SubMenuItems `json:"subMenuItems"`
	Title          string         `json:"title"`
	Icon           Icon           `json:"icon"`
	Accessibility  Accessibility  `json:"accessibility"`
	Tooltip        string         `json:"tooltip"`
	TrackingParams string         `json:"trackingParams"`
}
type SortMenu struct {
	SortFilterSubMenuRenderer SortFilterSubMenuRenderer `json:"sortFilterSubMenuRenderer"`
}
type TitleText struct {
	Runs []Runs `json:"runs"`
}
type CommentsCount struct {
	Runs []Runs `json:"runs"`
}
type CommentsHeaderRenderer struct {
	CountText         CountText         `json:"countText"`
	CreateRenderer    CreateRenderer    `json:"createRenderer"`
	SortMenu          SortMenu          `json:"sortMenu"`
	TrackingParams    string            `json:"trackingParams"`
	TitleText         TitleText         `json:"titleText"`
	CommentsCount     CommentsCount     `json:"commentsCount"`
	ShowSeparator     bool              `json:"showSeparator"`
	LoggingDirectives LoggingDirectives `json:"loggingDirectives"`
}
type Header struct {
	CommentsHeaderRenderer CommentsHeaderRenderer `json:"commentsHeaderRenderer"`
}
type ItemSectionContinuation struct {
	Contents []struct {
		CommentThreadRenderer CommentThreadRenderer `json:"commentThreadRenderer,omitempty"`
	} `json:"contents"`
	Continuations     []Continuations   `json:"continuations"`
	TrackingParams    string            `json:"trackingParams"`
	Header            Header            `json:"header"`
	SectionIdentifier string            `json:"sectionIdentifier"`
	LoggingDirectives LoggingDirectives `json:"loggingDirectives"`
}
type ContinuationContents struct {
	ItemSectionContinuation ItemSectionContinuation `json:"itemSectionContinuation"`
}
type Response struct {
	ResponseContext      ResponseContext      `json:"responseContext"`
	ContinuationContents ContinuationContents `json:"continuationContents"`
	TrackingParams       string               `json:"trackingParams"`
}
type Info struct {
	St int `json:"st"`
}
type Timing struct {
	Info Info `json:"info"`
}
