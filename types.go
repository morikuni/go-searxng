package searxng

import (
	"encoding/json"
	"fmt"
)

// Category is a category of search engines.
type Category string

var AllCategories = []Category{
	CategoryGeneral,
	CategoryImages,
	CategoryVideos,
	CategoryNews,
	CategoryMap,
	CategoryMusic,
	CategoryIT,
	CategoryScience,
	CategoryFiles,
	CategorySocialMedia,
}

const (
	CategoryGeneral     Category = "general"
	CategoryImages      Category = "images"
	CategoryVideos      Category = "videos"
	CategoryNews        Category = "news"
	CategoryMap         Category = "map"
	CategoryMusic       Category = "music"
	CategoryIT          Category = "it"
	CategoryScience     Category = "science"
	CategoryFiles       Category = "files"
	CategorySocialMedia Category = "social media"
)

type SearchEngine string

const (
	SearchEngineBing                  SearchEngine = "bing"
	SearchEngineBrave                 SearchEngine = "brave"
	SearchEngineDuckDuckGo            SearchEngine = "duckduckgo"
	SearchEngineGigablast             SearchEngine = "gigablast"
	SearchEngineGoogle                SearchEngine = "google"
	SearchEngineMojeek                SearchEngine = "mojeek"
	SearchEngineQwant                 SearchEngine = "qwant"
	SearchEngineStartpage             SearchEngine = "startpage"
	SearchEngineWiby                  SearchEngine = "wiby"
	SearchEngineYahoo                 SearchEngine = "yahoo"
	SearchEngineSeznam                SearchEngine = "seznam"
	SearchEngineGoo                   SearchEngine = "goo"
	SearchEngineNaver                 SearchEngine = "naver"
	SearchEngineAlexandria            SearchEngine = "alexandria"
	SearchEngineArchiveIs             SearchEngine = "archive is"
	SearchEngineCurlie                SearchEngine = "curlie"
	SearchEngineCurrency              SearchEngine = "currency"
	SearchEngineDdgDefinitions        SearchEngine = "ddg definitions"
	SearchEngineDictzone              SearchEngine = "dictzone"
	SearchEngineLingva                SearchEngine = "lingva"
	SearchEngineMarginalia            SearchEngine = "marginalia"
	SearchEnginePetalsearch           SearchEngine = "petalsearch"
	SearchEngineTineye                SearchEngine = "tineye"
	SearchEngineWikibooks             SearchEngine = "wikibooks"
	SearchEngineWikidata              SearchEngine = "wikidata"
	SearchEngineWikipedia             SearchEngine = "wikipedia"
	SearchEngineWikiquote             SearchEngine = "wikiquote"
	SearchEngineWikisource            SearchEngine = "wikisource"
	SearchEngineWikiversity           SearchEngine = "wikiversity"
	SearchEngineWikivoyage            SearchEngine = "wikivoyage"
	SearchEngineYep                   SearchEngine = "yep"
	SearchEngineWikimini              SearchEngine = "wikimini"
	SearchEngineBingImages            SearchEngine = "bing images"
	SearchEngineDuckDuckGoImages      SearchEngine = "duckduckgo images"
	SearchEngineGoogleImages          SearchEngine = "google images"
	SearchEngineQwantImages           SearchEngine = "qwant images"
	SearchEngine1x                    SearchEngine = "1x"
	SearchEngineArtic                 SearchEngine = "artic"
	SearchEngineDeviantart            SearchEngine = "deviantart"
	SearchEngineFlickr                SearchEngine = "flickr"
	SearchEngineFrinkiac              SearchEngine = "frinkiac"
	SearchEngineLibraryOfCongress     SearchEngine = "library of congress"
	SearchEngineOpenverse             SearchEngine = "openverse"
	SearchEnginePetalsearchImages     SearchEngine = "petalsearch images"
	SearchEngineUnsplash              SearchEngine = "unsplash"
	SearchEngineBingVideos            SearchEngine = "bing videos"
	SearchEngineGoogleVideos          SearchEngine = "google videos"
	SearchEngineQwantVideos           SearchEngine = "qwant videos"
	SearchEngineCccTv                 SearchEngine = "ccc-tv"
	SearchEngineDailymotion           SearchEngine = "dailymotion"
	SearchEngineGooglePlayMovies      SearchEngine = "google play movies"
	SearchEngineInvidious             SearchEngine = "invidious"
	SearchEnginePeertube              SearchEngine = "peertube"
	SearchEngineRumble                SearchEngine = "rumble"
	SearchEngineSepiasearch           SearchEngine = "sepiasearch"
	SearchEngineVimeo                 SearchEngine = "vimeo"
	SearchEngineYouTube               SearchEngine = "youtube"
	SearchEngineMediathekviewweb      SearchEngine = "mediathekviewweb"
	SearchEngineIna                   SearchEngine = "ina"
	SearchEngineBingNews              SearchEngine = "bing news"
	SearchEngineGoogleNews            SearchEngine = "google news"
	SearchEngineQwantNews             SearchEngine = "qwant news"
	SearchEngineWikinews              SearchEngine = "wikinews"
	SearchEngineYahooNews             SearchEngine = "yahoo news"
	SearchEngineAppleMaps             SearchEngine = "apple maps"
	SearchEngineOpenStreetMap         SearchEngine = "openstreetmap"
	SearchEnginePhoton                SearchEngine = "photon"
	SearchEngineAzlyrics              SearchEngine = "azlyrics"
	SearchEngineGenius                SearchEngine = "genius"
	SearchEngineBandcamp              SearchEngine = "bandcamp"
	SearchEngineDeezer                SearchEngine = "deezer"
	SearchEngineGpodder               SearchEngine = "gpodder"
	SearchEngineMixcloud              SearchEngine = "mixcloud"
	SearchEngineSoundcloud            SearchEngine = "soundcloud"
	SearchEngineBitbucket             SearchEngine = "bitbucket"
	SearchEngineCodeberg              SearchEngine = "codeberg"
	SearchEngineGithub                SearchEngine = "github"
	SearchEngineGitlab                SearchEngine = "gitlab"
	SearchEngineSourcehut             SearchEngine = "sourcehut"
	SearchEngineArchLinuxWiki         SearchEngine = "arch linux wiki"
	SearchEngineFreeSoftwareDirectory SearchEngine = "free software directory"
	SearchEngineGentoo                SearchEngine = "gentoo"
	SearchEngineFramalibre            SearchEngine = "framalibre"
	SearchEngineHabrahabr             SearchEngine = "habrahabr"
	SearchEngineLobsters              SearchEngine = "lobste.rs"
	SearchEngineMankier               SearchEngine = "mankier"
	SearchEngineSearchcodeCode        SearchEngine = "searchcode code"
	SearchEngineArxiv                 SearchEngine = "arxiv"
	SearchEngineCrossref              SearchEngine = "crossref"
	SearchEngineGoogleScholar         SearchEngine = "google scholar"
	SearchEnginePubmed                SearchEngine = "pubmed"
	SearchEngineSemanticScholar       SearchEngine = "semantic scholar"
	SearchEngineOpenAireDatasets      SearchEngine = "openairedatasets"
	SearchEngineOpenAirePublications  SearchEngine = "openairepublications"
	SearchEnginePdbe                  SearchEngine = "pdbe"
	SearchEngineApkMirror             SearchEngine = "apk mirror"
	SearchEngineAppleAppStore         SearchEngine = "apple app store"
	SearchEngineFdroid                SearchEngine = "fdroid"
	SearchEngineGooglePlayApps        SearchEngine = "google play apps"
	SearchEngine1337x                 SearchEngine = "1337x"
	SearchEngineBtdigg                SearchEngine = "btdigg"
	SearchEngineKickass               SearchEngine = "kickass"
	SearchEngineLibraryGenesis        SearchEngine = "library genesis"
	SearchEngineNyaa                  SearchEngine = "nyaa"
	SearchEngineOpenrepos             SearchEngine = "openrepos"
	SearchEnginePiratebay             SearchEngine = "piratebay"
	SearchEngineSolidtorrents         SearchEngine = "solidtorrents"
	SearchEngineTokyotoshokan         SearchEngine = "tokyotoshokan"
	SearchEngineAskUbuntu             SearchEngine = "askubuntu"
	SearchEngineStackOverflow         SearchEngine = "stackoverflow"
	SearchEngineSuperuser             SearchEngine = "superuser"
)

type TimeRange string

const (
	// TimeRangeDay is the 24 hours time range.
	TimeRangeDay TimeRange = "day"
	// TimeRangeWeek is the 1 week time range.
	TimeRangeWeek TimeRange = "week"
	// TimeRangeMonth is the 1 month time range.
	TimeRangeMonth TimeRange = "month"
	// TimeRangeYear is the 1 year time range.
	TimeRangeYear TimeRange = "year"
)

// SafeSearchLevel is the level of search result filter.
type SafeSearchLevel int

const (
	// SafeSearchLevelNone does not filter search results.
	SafeSearchLevelNone SafeSearchLevel = 0
	// SafeSearchLevelModerate filters search results moderately.
	SafeSearchLevelModerate SafeSearchLevel = 1
	// SafeSearchLevelStrict filters search results strictly.
	SafeSearchLevelStrict SafeSearchLevel = 2
)

type SearchResult interface {
	isSearchResult()
}

func newSearchResultByCategory(c Category) SearchResult {
	switch c {
	case CategoryGeneral:
		return &GeneralResult{}
	case CategoryImages:
		return &ImageResult{}
	case CategoryVideos:
		return &VideoResult{}
	case CategoryNews:
		return &NewsResult{}
	case CategoryMap:
		return &MapResult{}
	case CategoryMusic:
		return &MusicResult{}
	case CategoryIT:
		return &ITResult{}
	case CategoryScience:
		return &ScienceResult{}
	case CategoryFiles:
		return &FilesResult{}
	case CategorySocialMedia:
		return &SocialMediaResult{}
	default:
		panic(fmt.Sprintf("unknown category: %q", c))
	}
}

type GeneralResult struct {
	Category      Category       `json:"category"`
	Content       string         `json:"content"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate *string        `json:"publishedDate,omitempty"`
	IframeSrc     *string        `json:"iframe_src,omitempty"`
	Thumbnail     *string        `json:"thumbnail,omitempty"`
}

func (*GeneralResult) isSearchResult() {}

type ImageResult struct {
	Category     Category       `json:"category"`
	Engine       SearchEngine   `json:"engine"`
	Engines      []SearchEngine `json:"engines"`
	ParsedURL    []string       `json:"parsed_url"`
	Positions    []int          `json:"positions"`
	Score        float64        `json:"score"`
	Template     string         `json:"template"`
	Title        string         `json:"title"`
	URL          string         `json:"url"`
	ImgSrc       string         `json:"img_src"`
	Source       *string        `json:"source,omitempty"`
	ThumbnailSrc *string        `json:"thumbnail_src,omitempty"`
	Content      *string        `json:"content,omitempty"`
	Author       *string        `json:"author,omitempty"`
}

func (*ImageResult) isSearchResult() {}

type VideoResult struct {
	Category      Category        `json:"category"`
	Content       string          `json:"content,omitempty"`
	Engine        SearchEngine    `json:"engine"`
	Engines       []SearchEngine  `json:"engines"`
	ParsedURL     []string        `json:"parsed_url"`
	Positions     []int           `json:"positions"`
	Score         float64         `json:"score"`
	Template      string          `json:"template"`
	Thumbnail     string          `json:"thumbnail"`
	Title         string          `json:"title"`
	URL           string          `json:"url"`
	IframeSrc     *string         `json:"iframe_src,omitempty"`
	Length        *StringOrNumber `json:"length,omitempty"`
	PublishedDate *string         `json:"publishedDate,omitempty"`
	Author        *string         `json:"author,omitempty"`
	Metadata      *string         `json:"metadata,omitempty"`
	Duration      *string         `json:"duration,omitempty"`
}

func (*VideoResult) isSearchResult() {}

type NewsResult struct {
	Category      Category       `json:"category"`
	Content       string         `json:"content"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate string         `json:"publishedDate"`
	Thumbnail     string         `json:"thumbnail"`
	Metadata      string         `json:"metadata"`
}

func (*NewsResult) isSearchResult() {}

type MapResult struct {
	Address       *MapAddress     `json:"address,omitempty"`
	AddressLabel  *string         `json:"address_label,omitempty"`
	BoundingBox   []json.Number   `json:"boundingbox"`
	Category      Category        `json:"category"`
	Content       string          `json:"content"`
	Data          []KeyLabelValue `json:"data,omitempty"`
	Engine        SearchEngine    `json:"engine"`
	Engines       []SearchEngine  `json:"engines"`
	Geojson       Geojson         `json:"geojson"`
	Latitude      json.Number     `json:"latitude"`
	Links         []LabeledURL    `json:"links,omitempty"`
	Longitude     json.Number     `json:"longitude"`
	OpenStreetMap OpenStreetMap   `json:"osm"`
	ParsedURL     []string        `json:"parsed_url"`
	Positions     []int           `json:"positions"`
	Score         float64         `json:"score"`
	Template      string          `json:"template"`
	Thumbnail     *string         `json:"thumbnail,omitempty"`
	Title         string          `json:"title"`
	Type          *string         `json:"type,omitempty"`
	TypeIcon      any             `json:"type_icon"` // I could not determine the type because all data is null.
	URL           string          `json:"url"`
}

func (*MapResult) isSearchResult() {}

type MapAddress struct {
	Country     string  `json:"country"`
	HouseNumber *string `json:"house_number,omitempty"`
	Locality    string  `json:"locality"`
	Name        string  `json:"name"`
	Postcode    *string `json:"postcode,omitempty"`
	Road        *string `json:"road,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
}

type Geojson struct {
	Coordinates PointOrPolygon `json:"coordinates"`
	Type        string         `json:"type"`
}

// PointOrPolygon is a either a Point([]json.Number) or a Polygon([][]json.Number).
type PointOrPolygon = any

type KeyLabelValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type LabeledURL struct {
	Label    string `json:"label"`
	URL      string `json:"url"`
	URLLabel string `json:"url_label"`
}

type OpenStreetMap struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

type MusicResult struct {
	Category      Category       `json:"category"`
	Content       *string        `json:"content,omitempty"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	IframeSrc     string         `json:"iframe_src"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Thumbnail     string         `json:"thumbnail"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate *string        `json:"publishedDate,omitempty"`
	Length        *json.Number   `json:"length,omitempty"`
	Author        *string        `json:"author,omitempty"`
}

func (*MusicResult) isSearchResult() {}

type ITResult struct {
	Category      Category       `json:"category"`
	Content       *string        `json:"content,omitempty"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Thumbnail     *string        `json:"thumbnail,omitempty"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate *string        `json:"publishedDate,omitempty"`
}

func (*ITResult) isSearchResult() {}

type ScienceResult struct {
	Category      Category       `json:"category"`
	Content       *string        `json:"content,omitempty"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate *string        `json:"publishedDate,omitempty"`
}

func (*ScienceResult) isSearchResult() {}

type FilesResult struct {
	Category      Category       `json:"category"`
	Content       *string        `json:"content,omitempty"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate *string        `json:"publishedDate,omitempty"`
	Metadata      *string        `json:"metadata,omitempty"`
	Thumbnail     *string        `json:"thumbnail,omitempty"`
}

func (*FilesResult) isSearchResult() {}

type SocialMediaResult struct {
	Category      Category       `json:"category"`
	Content       string         `json:"content"`
	Engine        SearchEngine   `json:"engine"`
	Engines       []SearchEngine `json:"engines"`
	ParsedURL     []string       `json:"parsed_url"`
	Positions     []int          `json:"positions"`
	Score         float64        `json:"score"`
	Template      string         `json:"template"`
	Title         string         `json:"title"`
	URL           string         `json:"url"`
	PublishedDate *string        `json:"publishedDate,omitempty"`
	Thumbnail     *string        `json:"thumbnail,omitempty"`
}

func (*SocialMediaResult) isSearchResult() {}

type SearchOutput struct {
	NumberOfResults     int            `json:"number_of_results"`
	Query               string         `json:"query"`
	Results             []SearchResult `json:"results"`
	UnresponsiveEngines [][2]string    `json:"unresponsive_engines"`
}

func (o *SearchOutput) UnmarshalJSON(data []byte) error {
	var result struct {
		NumberOfResults     int               `json:"number_of_results"`
		Query               string            `json:"query"`
		Results             []json.RawMessage `json:"results"`
		UnresponsiveEngines [][2]string       `json:"unresponsive_engines"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	// Results をデコード
	for _, result := range result.Results {
		// Category を一時的に取得
		var temp struct {
			Category Category `json:"category"`
		}
		if err := json.Unmarshal(result, &temp); err != nil {
			return err
		}

		// Category に基づいて型を決定
		var searchResult SearchResult
		switch temp.Category {
		case CategoryGeneral:
			searchResult = &GeneralResult{}
		case CategoryImages:
			searchResult = &ImageResult{}
		case CategoryVideos:
			searchResult = &VideoResult{}
		case CategoryNews:
			searchResult = &NewsResult{}
		case CategoryMap:
			searchResult = &MapResult{}
		case CategoryMusic:
			searchResult = &MusicResult{}
		case CategoryIT:
			searchResult = &ITResult{}
		case CategoryScience:
			searchResult = &ScienceResult{}
		case CategoryFiles:
			searchResult = &FilesResult{}
		case CategorySocialMedia:
			searchResult = &SocialMediaResult{}
		default:
			return fmt.Errorf("unknown category: %q", temp.Category)
		}

		// 適切な型にデコード
		if err := json.Unmarshal(result, searchResult); err != nil {
			return err
		}

		o.Results = append(o.Results, searchResult)
	}

	return nil
}

type StringOrNumber struct {
	String *string
	Number *json.Number
}

func (s *StringOrNumber) UnmarshalJSON(data []byte) error {
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		s.Number = &n
		return nil
	}
	return json.Unmarshal(data, &s.String)
}

func (s *StringOrNumber) MarshalJSON() ([]byte, error) {
	if s.String != nil {
		return json.Marshal(s.String)
	}
	return json.Marshal(s.Number)
}

func (s *StringOrNumber) IsString() bool {
	return s.String != nil
}

func (s *StringOrNumber) IsNumber() bool {
	return s.Number != nil
}

func (s *StringOrNumber) StringValue() string {
	if s.String != nil {
		return *s.String
	}
	return s.Number.String()
}

func (s *StringOrNumber) NumberValue() (number json.Number, valid bool) {
	if s.Number != nil {
		return *s.Number, true
	}
	return "", false
}
