package client

type Competition struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type Country struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha_2"`
	Alpha3 string `json:"alpha_3"`
}

type Event struct {
	ID          int            `json:"id"`
	DateUTC     string         `json:"date_utc"`
	TimeUTC     string         `json:"time_utc"`
	DateTime    string         `json:"date_time"`
	TeamA       Team           `json:"team_A"`
	TeamB       Team           `json:"team_B"`
	Tournament  Tournament     `json:"tournament"`
	Stage       Stage          `json:"stage,omitempty"`
	Week        Week           `json:"week,omitempty"`
	HTS_A       int            `json:"hts_A"`
	HTS_B       int            `json:"hts_B"`
	FS_A        int            `json:"fs_A"`
	FS_B        int            `json:"fs_B"`
	Total_A     int            `json:"total_A"`
	Total_B     int            `json:"total_B"`
	MatchLength string         `json:"match_length"`
	MatchPeriod int            `json:"match_period"`
	Status      string         `json:"status"`
	Coverage    string         `json:"coverage,omitempty"`
	Period      []Period       `json:"period,omitempty"`
	Venue       Venue          `json:"venue"`
	TVChannel   []TVChannel    `json:"tv_channel,omitempty"`
	Occurrence  []Occurrence_s `json:"occurrence,omitempty"`
}

type Lineup struct {
	TeamAManager Person        `json:"team_A_manager"`
	TeamALineup  []SquadMember `json:"team_A_lineup"`
	TeamBManager Person        `json:"team_B_manager"`
	TeamBLineup  []SquadMember `json:"team_B_lineup"`
}

type Platform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Media_s struct {
	ID          int      `json:"id"`
	ContentType string   `json:"content_type"`
	EmbedCode   string   `json:"embed"`
	Created     string   `json:"created"`
	Modified    string   `json:"modified"`
	Platform    Platform `json:"platform"`
}

type Occurrence_s struct {
	ID           int       `json:"id"`
	MatchPeriod  int       `json:"match_period"`
	Minute       int       `json:"minute"`
	MinuteExtra  int       `json:"minute_extra,omitempty"`
	In           string    `json:"in,omitempty"`
	Out          string    `json:"out,omitempty"`
	Team         Team      `json:"team,omitempty"`
	Player       Person    `json:"player,omitempty"`
	PlayerOff    Person    `json:"player_off,omitempty"`
	Reason       string    `json:"reason,omitempty"`
	AssistPlayer Person    `json:"assist_player,omitempty"`
	TypeCode     string    `json:"type_code"`
	TypeName     string    `json:"type_name"`
	Text         string    `json:"text,omitempty"`
	Media        []Media_s `json:"media,omitempty"`
	TeamAScore   *int      `json:"team_A_score,omitempty"`
	TeamBScore   *int      `json:"team_B_score,omitempty"`
	VarType      string    `json:"var_type,omitempty"`
	VarDecision  string    `json:"var_decision,omitempty"`
	Outcome      string    `json:"outcome,omitempty"`
}

type OccurrenceResponse = []Occurrence_s

type Period struct {
	Period int    `json:"period"`
	Start  string `json:"start"`
	End    string `json:"end"`
}

type Person struct {
	ID          int     `json:"id,omitempty"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	MatchName   string  `json:"match_name,omitempty"`
	Type        string  `json:"type,omitempty"`
	Position    string  `json:"position,omitempty"`
	Photo       string  `json:"photo,omitempty"`
	Height      int     `json:"height,omitempty"`
	Weight      int     `json:"weight,omitempty"`
	BirthDate   string  `json:"birth_date,omitempty"`
	BirthPlace  string  `json:"birth_place,omitempty"`
	Nationality Country `json:"nationality,omitempty"`
}

type Squad struct {
	ID    int           `json:"id"`
	Team  Team          `json:"team"`
	Squad []SquadMember `json:"squad"`
}

type SquadMember struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MatchName   string `json:"match_name"`
	ShirtNumber int    `json:"shirt_number,omitempty"`
	Position    string `json:"position,omitempty"`
	Number      int    `json:"number,omitempty"`
	Photo       string `json:"photo,omitempty"`
	Substitute  bool   `json:"substitute,omitempty"`
}

type Stage struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	StartDate    string          `json:"start_date"`
	EndDate      string          `json:"end_date"`
	HasStandings bool            `json:"has_standings,omitempty"`
	Standings    []StandingEntry `json:"standings,omitempty"`
}

type StandingEntry struct {
	Position       int  `json:"position"`
	LastPosition   int  `json:"last_position"`
	Points         int  `json:"points"`
	Played         int  `json:"played"`
	Won            int  `json:"won"`
	Drawn          int  `json:"drawn"`
	Lost           int  `json:"lost"`
	GoalsFor       int  `json:"goals_for"`
	GoalsAgainst   int  `json:"goals_against"`
	GoalDifference int  `json:"goal_difference"`
	Team           Team `json:"team"`
}

type Standings struct {
	TournamentID int         `json:"id"`
	Name         string      `json:"name"`
	StartDate    string      `json:"start_date"`
	EndDate      string      `json:"end_date"`
	Season       string      `json:"season"`
	Competition  Competition `json:"competition"`
	Area         Country     `json:"area"`
	Stage        []Stage     `json:"stage"`
}

type Stats struct {
	Played          int `json:"played"`
	Won             int `json:"won"`
	Drawn           int `json:"drawn"`
	Lost            int `json:"lost"`
	GoalsFor        int `json:"goals_for"`
	GoalsAgainst    int `json:"goals_against"`
	GoalsDifference int `json:"goals_difference"`
	Points          int `json:"points"`
}

type Team struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	OfficialName string  `json:"official_name,omitempty"`
	Code         string  `json:"code,omitempty"`
	Type         string  `json:"type,omitempty"`
	Gender       string  `json:"gender"`
	City         string  `json:"city,omitempty"`
	Country      Country `json:"country,omitempty"`
	Logo         string  `json:"logo"`
}

type TeamDetailed struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	OfficialName string        `json:"official_name"`
	Code         string        `json:"code"`
	Type         string        `json:"type"`
	Gender       string        `json:"gender"`
	City         string        `json:"city"`
	Country      Country       `json:"country"`
	Logo         string        `json:"logo"`
	Lineup       []SquadMember `json:"lineup"`
	Referee      Person        `json:"referee"`
	TVChannel    string        `json:"tv_channel"`
}

type Tournament struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Active      bool        `json:"active"`
	StartDate   string      `json:"start_date"`
	EndDate     string      `json:"end_date"`
	Season      string      `json:"season"`
	Competition Competition `json:"competition"`
	Area        Country     `json:"area"`
}

type TVChannel struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	Country Country `json:"country"`
}

type Venue struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Country Country `json:"country"`
	Photo   string  `json:"photo"`
}

type Week struct {
	Index     int    `json:"index"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
