package service

import (
	"encoding/xml"
)

type TaskInfo struct {
	Name   []string
	Status bool // is show
}

type StartChan struct {
	Start chan bool
	IsClosed bool
}

type Page struct {
	XMLName           xml.Name `xml:"Page"`
	Text              string   `xml:",chardata"`
	One               string   `xml:"one,attr"`
	ID                string   `xml:"ID,attr"`
	Name              string   `xml:"name,attr"`
	DateTime          string   `xml:"dateTime,attr"`
	LastModifiedTime  string   `xml:"lastModifiedTime,attr"`
	PageLevel         string   `xml:"pageLevel,attr"`
	IsCurrentlyViewed string   `xml:"isCurrentlyViewed,attr"`
	Lang              string   `xml:"lang,attr"`
	TagDef            []struct {
		Text           string `xml:",chardata"`
		Index          string `xml:"index,attr"`
		Type           string `xml:"type,attr"`
		Symbol         string `xml:"symbol,attr"`
		FontColor      string `xml:"fontColor,attr"`
		HighlightColor string `xml:"highlightColor,attr"`
		Name           string `xml:"name,attr"`
	} `xml:"TagDef"`
	QuickStyleDef []struct {
		Text           string `xml:",chardata"`
		Index          string `xml:"index,attr"`
		Name           string `xml:"name,attr"`
		FontColor      string `xml:"fontColor,attr"`
		HighlightColor string `xml:"highlightColor,attr"`
		Font           string `xml:"font,attr"`
		FontSize       string `xml:"fontSize,attr"`
		SpaceBefore    string `xml:"spaceBefore,attr"`
		SpaceAfter     string `xml:"spaceAfter,attr"`
	} `xml:"QuickStyleDef"`
	PageSettings struct {
		Text     string `xml:",chardata"`
		RTL      string `xml:"RTL,attr"`
		Color    string `xml:"color,attr"`
		PageSize struct {
			Text      string `xml:",chardata"`
			Automatic string `xml:"Automatic"`
		} `xml:"PageSize"`
		RuleLines struct {
			Text    string `xml:",chardata"`
			Visible string `xml:"visible,attr"`
		} `xml:"RuleLines"`
	} `xml:"PageSettings"`
	Title struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		OE   struct {
			Text                       string `xml:",chardata"`
			Author                     string `xml:"author,attr"`
			AuthorInitials             string `xml:"authorInitials,attr"`
			AuthorResolutionID         string `xml:"authorResolutionID,attr"`
			LastModifiedBy             string `xml:"lastModifiedBy,attr"`
			LastModifiedByInitials     string `xml:"lastModifiedByInitials,attr"`
			LastModifiedByResolutionID string `xml:"lastModifiedByResolutionID,attr"`
			CreationTime               string `xml:"creationTime,attr"`
			LastModifiedTime           string `xml:"lastModifiedTime,attr"`
			ObjectID                   string `xml:"objectID,attr"`
			Alignment                  string `xml:"alignment,attr"`
			QuickStyleIndex            string `xml:"quickStyleIndex,attr"`
			T                          string `xml:"T"`
		} `xml:"OE"`
	} `xml:"Title"`
	Outline []struct {
		Text                   string `xml:",chardata"`
		Author                 string `xml:"author,attr"`
		AuthorInitials         string `xml:"authorInitials,attr"`
		LastModifiedBy         string `xml:"lastModifiedBy,attr"`
		LastModifiedByInitials string `xml:"lastModifiedByInitials,attr"`
		LastModifiedTime       string `xml:"lastModifiedTime,attr"`
		ObjectID               string `xml:"objectID,attr"`
		Lang                   string `xml:"lang,attr"`
		Position               struct {
			Text string `xml:",chardata"`
			X    string `xml:"x,attr"`
			Y    string `xml:"y,attr"`
			Z    string `xml:"z,attr"`
		} `xml:"Position"`
		Size struct {
			Text        string `xml:",chardata"`
			Width       string `xml:"width,attr"`
			Height      string `xml:"height,attr"`
			IsSetByUser string `xml:"isSetByUser,attr"`
		} `xml:"Size"`
		OEChildren struct {
			Text string `xml:",chardata"`
			OE   []struct {
				Text                       string `xml:",chardata"`
				LastModifiedByResolutionID string `xml:"lastModifiedByResolutionID,attr"`
				CreationTime               string `xml:"creationTime,attr"`
				LastModifiedTime           string `xml:"lastModifiedTime,attr"`
				ObjectID                   string `xml:"objectID,attr"`
				Alignment                  string `xml:"alignment,attr"`
				QuickStyleIndex            string `xml:"quickStyleIndex,attr"`
				Style                      string `xml:"style,attr"`
				AuthorInitials             string `xml:"authorInitials,attr"`
				AuthorResolutionID         string `xml:"authorResolutionID,attr"`
				Lang                       string `xml:"lang,attr"`
				T                          struct {
					Text  string `xml:",chardata"`
					Style string `xml:"style,attr"`
				} `xml:"T"`
				Tag struct {
					Text           string `xml:",chardata"`
					Index          string `xml:"index,attr"`
					Completed      string `xml:"completed,attr"`
					Disabled       string `xml:"disabled,attr"`
					CreationDate   string `xml:"creationDate,attr"`
					CompletionDate string `xml:"completionDate,attr"`
				} `xml:"Tag"`
				OEChildren struct {
					Text string `xml:",chardata"`
					OE   []struct {
						Text                       string `xml:",chardata"`
						AuthorResolutionID         string `xml:"authorResolutionID,attr"`
						LastModifiedByResolutionID string `xml:"lastModifiedByResolutionID,attr"`
						CreationTime               string `xml:"creationTime,attr"`
						LastModifiedTime           string `xml:"lastModifiedTime,attr"`
						ObjectID                   string `xml:"objectID,attr"`
						Alignment                  string `xml:"alignment,attr"`
						QuickStyleIndex            string `xml:"quickStyleIndex,attr"`
						Style                      string `xml:"style,attr"`
						List                       struct {
							Text   string `xml:",chardata"`
							Bullet struct {
								Text     string `xml:",chardata"`
								Bullet   string `xml:"bullet,attr"`
								FontSize string `xml:"fontSize,attr"`
							} `xml:"Bullet"`
						} `xml:"List"`
						T string `xml:"T"`
					} `xml:"OE"`
				} `xml:"OEChildren"`
				Table struct {
					Text             string `xml:",chardata"`
					BordersVisible   string `xml:"bordersVisible,attr"`
					HasHeaderRow     string `xml:"hasHeaderRow,attr"`
					LastModifiedTime string `xml:"lastModifiedTime,attr"`
					ObjectID         string `xml:"objectID,attr"`
					Columns          struct {
						Text   string `xml:",chardata"`
						Column struct {
							Text  string `xml:",chardata"`
							Index string `xml:"index,attr"`
							Width string `xml:"width,attr"`
						} `xml:"Column"`
					} `xml:"Columns"`
					Row []struct {
						Text             string `xml:",chardata"`
						ObjectID         string `xml:"objectID,attr"`
						LastModifiedTime string `xml:"lastModifiedTime,attr"`
						Cell             struct {
							Text                   string `xml:",chardata"`
							ShadingColor           string `xml:"shadingColor,attr"`
							LastModifiedTime       string `xml:"lastModifiedTime,attr"`
							ObjectID               string `xml:"objectID,attr"`
							LastModifiedByInitials string `xml:"lastModifiedByInitials,attr"`
							OEChildren             struct {
								Text string `xml:",chardata"`
								OE   []struct {
									Text                       string `xml:",chardata"`
									AuthorResolutionID         string `xml:"authorResolutionID,attr"`
									LastModifiedByResolutionID string `xml:"lastModifiedByResolutionID,attr"`
									CreationTime               string `xml:"creationTime,attr"`
									LastModifiedTime           string `xml:"lastModifiedTime,attr"`
									ObjectID                   string `xml:"objectID,attr"`
									Alignment                  string `xml:"alignment,attr"`
									QuickStyleIndex            string `xml:"quickStyleIndex,attr"`
									Style                      string `xml:"style,attr"`
									Lang                       string `xml:"lang,attr"`
									T                          string `xml:"T"`
								} `xml:"OE"`
							} `xml:"OEChildren"`
						} `xml:"Cell"`
					} `xml:"Row"`
				} `xml:"Table"`
			} `xml:"OE"`
		} `xml:"OEChildren"`
	} `xml:"Outline"`
}
