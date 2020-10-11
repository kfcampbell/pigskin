package responses

// this package holds models for fleaflicker responses

// LeagueScoreboard returns all of the games for the current week
type LeagueScoreboard struct {
	SchedulePeriod          ScoringPeriodRange   `json:"schedulePeriod"`
	EligibleSchedulePeriods []ScoringPeriodRange `json:"eligibleSchedulePeriods"`
	Games                   []FantasyGame        `json:"games"`
	InProgress              bool                 `json:"isInProgress"`
}

// FantasyGame describes a fantasy game
type FantasyGame struct {
	ID           string           `json:"id"`
	Away         Team             `json:"away"`
	Home         Team             `json:"home"`
	AwayScore    FantasyLineScore `json:"awayScore"`
	HomeScore    FantasyLineScore `json:"homeScore"`
	HomeResult   string           `json:"homeResult"`
	AwayResult   string           `json:"awayResult"`
	IsInProgress bool             `json:"isInProgress"`
	IsFinalScore bool             `json:"isFinalScore"`
}

// Team describes a fantasy team
type Team struct {
	ID                      int32            `json:"id"`
	Name                    string           `json:"name"`
	LogoURL                 string           `json:"logoUrl"`
	OverallRecord           TeamRecord       `json:"recordOverall"`
	DivisionRecord          TeamRecord       `json:"recordDivision"`
	PostseasonRecord        TeamRecord       `json:"recordPostseason"`
	PointsFor               FormattedDecimal `json:"pointsFor"`
	PointsAgainst           FormattedDecimal `json:"pointsAgainst"`
	Streak                  FormattedDecimal `json:"streak"`
	WaiverPosition          int              `json:"waiverPosition"`
	WaiverAcquisitionBudget FormattedDecimal `json:"waiverAcquisitionBudget"`
	DraftPosition           int              `json:"draftPosition"`
	Owners                  []User           `json:"owners"`
}

// TeamRecord describes a record that belongs to a team
type TeamRecord struct {
	Wins          int              `json:"wins"`
	Losses        int              `json:"losses"`
	Ties          int              `json:"ties"`
	WinPercentage FormattedDecimal `json:"winPercentage"`
	Rank          int              `json:"rank"`
	Formatted     string           `json:"formatted"`
}

// FantasyLineScore describes an in-progress fantasy score
type FantasyLineScore struct {
	YetToPlay              int              `json:"yetToPlay"`
	YetToPlayPositions     []string         `json:"yetToPlayPositions"`
	InPlay                 int32            `json:"inPlay"`
	InPlayPositions        []string         `json:"inPlayPositions"`
	AlreadyPlayed          int32            `json:"alreadyPlayed"`
	AlreadyPlayedPositions []string         `json:"alreadyPlayedPositions"`
	Score                  FormattedDecimal `json:"score"`
	ProjectedScore         FormattedDecimal `json:"projected"`
}

// LeagueBoxscore shows the actual, optimum, and projected scores for a given game
type LeagueBoxscore struct {
	Game       FantasyGame        `json:"game"`
	Lineups    MatchupRosterGroup `json:"matchupRosterGroup"`
	PointsAway MatchupTeamScore   `json:"pointsAway"`
	PointsHome MatchupTeamScore   `json:"pointsHome"`
	// one, two, skip a few
	IsInProgress bool `json:"isInProgress"`
}

// MatchupRosterGroup represents the lineups in a fantasy game
type MatchupRosterGroup struct {
	Group string `json:"group"`
}

// MatchupTeamScore shows a team's score and their win probability
type MatchupTeamScore struct {
	Total               MatchupTeamSum `json:"total"`
	ScoringPeriod       ScoringPeriod  `json:"scoringPeriod"`
	WinProbability      float32        `json:"winProbability"`
	IsWinProbabilitySet bool           `json:"isWinProbabilitySet"`
}

// MatchupTeamSum contains the team score, the optimum score, and the projected score.
type MatchupTeamSum struct {
	Value     FormattedDecimal `json:"value"`
	Optimum   FormattedDecimal `json:"optimum"`
	Projected FormattedDecimal `json:"projected"`
	// Then the actual optimum lineup is stored here
}

// MatchupRosterSlot represents the slot for a player in a particular matchup
type MatchupRosterSlot struct {
	Position RosterPosition `json:"position"`
}

// ProPlayer represents a professional player
type ProPlayer struct {
	Sport           string  `json:"sport"`
	ID              int32   `json:"id"`
	FullName        string  `json:"nameFull"`
	ShortName       string  `json:"nameShort"`
	FirstName       string  `json:"nameFirst"`
	LastName        string  `json:"nameLast"`
	ProTeamAbbrev   string  `json:"proTeamAbbrevation"`
	ProTeam         ProTeam `json:"proTeam"`
	Position        string  `json:"position"`
	HeadshotURL     string  `json:"headshotUrl"`
	IsTeamAggregate bool    `json:"isTeamAggregate"`
	NFLByeWeek      int32   `json:"nflByeWeek"`
	Injury          Injury  `json:"injury"`
}

// Injury represent's a single ProPlayer's injury
type Injury struct {
	// LOL at this misspelling
	TypeAbbreviation string `json:"typeAbbreviaition"`
	Description      string `json:"description"`
	Severity         string `json:"severity"`
	FullType         string `json:"typeFull"`
}

// ProTeam represents a professional team
type ProTeam struct {
	Abbrev    string `json:"abbreviation"`
	Location  string `json:"location"`
	Name      string `json:"name"`
	FreeAgent bool   `json:"is_free_agent"`
}

// RosterPosition represents a player's position on the roster
type RosterPosition struct {
	Label       string   `json:"label"`
	Group       string   `json:"group"`
	Eligibility []string `json:"eligibility"`
	Min         int32    `json:"min"`
	Max         int32    `json:"max"`
	Start       int32    `json:"start"`
	Colors      []string `json:"colors"`
}

// LeaguePlayer represents a player in a given week
type LeaguePlayer struct {
	ProPlayer ProPlayer `json:"proPlayer"`
	// Bunch of super interesting player stats points stuff below here
}

// LeaguePlayerGame represents a player's in-game performance in a given week
type LeaguePlayerGame struct {
	Participant string        `json:"participant"` // HOME or AWAY
	Period      ScoringPeriod `json:"scoringPeriod"`
	Game        ProGame       `json:"proGame"`
	// More stuff I'm lazy
}

// ProGame is a placeholder. i'm getting somewhat lazy
type ProGame struct {
}

// FormattedDecimal is a lame way to format decimals
type FormattedDecimal struct {
	Value     float32 `json:"value"`
	Formatted string  `json:"formatted"`
}

// User is the owner of a fantasy team
type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	LastSeen    int64  `json:"last_seen"`
	Initials    string `json:"initials"`
	LastSeenISO string `json:"last_seen_iso"`
}

// ScoringPeriodRange shows a scoring period range
type ScoringPeriodRange struct {
	Ordinal int           `json:"ordinal"`
	Low     ScoringPeriod `json:"low"`
	Value   int           `json:"value"`
	//High        ScoringPeriod `json:"high"`
	//ContainsNow bool          `json:"contains_now"`
}

// ScoringPeriod shows the duration of a scoring period
type ScoringPeriod struct {
	Ordinal      int    `json:"ordinal"`
	StartEpochMS string `json:"startEpochMilli"`
	//Duration     string `json:"duration"`
	//Season       int32  `json:"season"`
	//IsNow        bool   `json:"is_now"`
	//Label        string `json:"label"`
}

// Go has weak enum support, so...

// WinGameResult is a win
const WinGameResult = "WIN"

// LoseGameResult is a loss
const LoseGameResult = "LOSE"

// TieGameResult is a tie
const TieGameResult = "TIE"

// UndecidedGameResult is for in-progress games
const UndecidedGameResult = "UNDECIDED"
