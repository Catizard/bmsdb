package bmsdb

type LR2Score struct {
	MD5        string `gorm:"column:hash"`
	Clear      int    `gorm:"column:clear"`
	Perfect    int    `gorm:"column:perfect"`
	Great      int    `gorm:"column:great"`
	Good       int    `gorm:"column:good"`
	Bad        int    `gorm:"column:bad"`
	Poor       int    `gorm:"column:Poor"`
	TotalNotes int    `gorm:"column:totalnotes"`
	MaxCombo   int    `gorm:"column:maxcombo"`
	Minbp      int    `gorm:"column:minbp"`
	PlayCount  int    `gorm:"column:playcount"`
	ClearCount int    `gorm:"column:clearcount"`
	FailCount  int    `gorm:"column:failcount"`
	Rank       int    `gorm:"column:rank"`
	Rate       int    `gorm:"column:rate"`
	ClearDB    int    `gorm:"column:clear_db"`
	OpHistory  int    `gorm:"column:op_history"`
	ScoreHash  string `gorm:"column:scorehash"`
	Ghost      string `gorm:"column:ghost"`
	ClearSD    int    `gorm:"column:clear_sd"`
	ClearEX    int    `gorm:"column:clear_ex"`
	OpBest     int    `gorm:"column:op_best"`
	RSeed      int    `gorm:"column:rseed"`
	Complete   int    `gorm:"column:complete"`
	RowID      int    `gorm:"column:row_id"`
}

func (LR2Score) TableName() string {
	return "score"
}

type LR2Song struct {
	MD5        string `gorm:"column:hash"`
	Title      string `gorm:"column:title"`
	SubTitle   string `gorm:"column:subtitle"`
	Genre      string
	Artist     string
	SubArtist  string `gorm:"column:subartist"`
	Tag        string
	Path       string
	Folder     string
	Type       int    `gorm:"column:type"`
	StageFile  string `gorm:"column:stagefile"`
	Banner     string
	BackBmp    string `gorm:"column:backbmp"`
	Parent     string
	Level      int32
	Difficulty int32
	MaxBpm     int32 `gorm:"column:maxbpm"`
	MinBpm     int32 `gorm:"column:minbpm"`
	Length     int32
	Mode       int32
	Judge      int32
	LongNote   int   `gorm:"column:longnote"`
	Bga        int   `gorm:"column:bga"`
	Random     int   `gorm:"column:random"`
	Date       int64 `gorm:"column:date"`
	Favorite   int   `gorm:"column:favorite"`
	Txt        int   `gorm:"column:txt"`
	Karinotes  int   `gorm:"column:karinotes"`
	AddDate    int64 `gorm:"column:adddate"`
	ExLevel    int   `gorm:"exlevel"`
}

func (LR2Song) TableName() string {
	return "song"
}

type LR2ScanResult struct {
	Song   string
	Scores []LR2PlayerScoreFile
}

type LR2PlayerScoreFile struct {
	Name string
	Path string
}
