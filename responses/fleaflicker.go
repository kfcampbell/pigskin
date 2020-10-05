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
