package responses

// this package holds models for fleaflicker responses

// LeagueScoreboard returns all of the games for the current week
type LeagueScoreboard struct {
	Games      []FantasyGame `json:"games"`
	InProgress bool          `json:"is_in_progress"`
}

// FantasyGame describes a fantasy game
type FantasyGame struct {
	ID           int              `json:"id"`
	Away         Team             `json:"away"`
	Home         Team             `json:"home"`
	AwayScore    FantasyLineScore `json:"away_score"`
	HomeScore    FantasyLineScore `json:"home_score"`
	HomeResult   string           `json:"home_result"`
	AwayResult   string           `json:"away_result"`
	IsInProgress bool             `json:"is_in_progress"`
	IsFinalScore bool             `json:"is_final_score"`
}

// Team describes a fantasy team
type Team struct {
	ID                      int              `json:"id"`
	Name                    string           `json:"name"`
	LogoURL                 string           `json:"logo_url"`
	OverallRecord           TeamRecord       `json:"record_overall"`
	DivisionRecord          TeamRecord       `json:"record_division"`
	PostseasonRecord        TeamRecord       `json:"record_postseason"`
	PointsFor               FormattedDecimal `json:"points_for"`
	PointsAgainst           FormattedDecimal `json:"points_against"`
	Streak                  FormattedDecimal `json:"streak"`
	WaiverPosition          int              `json:"waver_position"`
	WaiverAcquisitionBudget FormattedDecimal `json:"waiver_acquisition_budget"`
	DraftPosition           int              `json:"draft_position"`
	Owners                  []User           `json:"owners"`
}

// TeamRecord describes a record that belongs to a team
type TeamRecord struct {
	Wins          int    `json:"wins"`
	Losses        int    `json:"losses"`
	Ties          int    `json:"ties"`
	WinPercentage int    `json:"win_percentage"`
	Rank          int    `json:"rank"`
	Formatted     string `json:"formatted"`
}

// FantasyLineScore describes an in-progress fantasy score
type FantasyLineScore struct {
	YetToPlay              int              `json:"yet_to_play"`
	YetToPlayPositions     []string         `json:"yet_to_play_positions"`
	InPlay                 int              `json:"in_play"`
	InPlayPositions        []string         `json:"in_play_positions"`
	AlreadyPlayed          int              `json:"already_played"`
	AlreadyPlayedPositions []string         `json:"already_played_positions"`
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
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	LastSeen    int64  `json:"last_seen"`
	Initials    string `json:"initials"`
	LastSeenISO string `json:"last_seen_iso"`
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
