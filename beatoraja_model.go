package bmsdb

type BeatorajaScoreLog struct {
	Sha256    string
	Mode      string
	Clear     int32
	OldClear  int32 `gorm:"column:oldclear"`
	Score     int32
	OldScore  int32 `gorm:"column:oldscore"`
	Combo     int32
	OldCombo  int32 `gorm:"column:oldcombo"`
	Minbp     int32
	OldMinbp  int32 `gorm:"column:oldminbp"`
	TimeStamp int64 `gorm:"column:date"`
}

func (BeatorajaScoreLog) TableName() string {
	return "scorelog"
}

type BeatorajaSongData struct {
	Md5        string
	Sha256     string
	Title      string
	SubTitle   string `gorm:"column:subtitle"`
	Genre      string
	Artist     string
	SubArtist  string `gorm:"column:subartist"`
	Tag        string
	Path       string
	Folder     string
	StageFile  string `gorm:"column:stagefile"`
	Banner     string
	BackBmp    string `gorm:"column:backbmp"`
	Preview    string
	Parent     string
	Level      int32
	Difficulty int32
	MaxBpm     int32 `gorm:"column:maxbpm"`
	MinBpm     int32 `gorm:"column:minbpm"`
	Length     int32
	Mode       int32
	Judge      int32
	Feature    int32
	Content    int32
	Date       int64
	Favorite   int32
	AddDate    int64 `gorm:"column:adddate"`
	Notes      int32
	ChartHash  string `gorm:"column:charthash"`
}

func (BeatorajaSongData) TableName() string {
	return "song"
}

type BeatorajaScoreDataLog struct {
	Sha256     string
	Mode       string
	Clear      int32
	Epg        int32
	Lpg        int32
	Egr        int32
	Lgr        int32
	Egd        int32
	Lgd        int32
	Ebd        int32
	Lbd        int32
	Epr        int32
	Lpr        int32
	Ems        int32
	Lms        int32
	Notes      int32
	Combo      int32
	Minbp      int32
	PlayCount  int32 `gorm:"column:playcount"`
	ClearCount int32 `gorm:"column:clearcount"`
	Option     int32
	Seed       int64
	Random     int32
	TimeStamp  int64 `gorm:"column:date"`
	State      int32
}

func (BeatorajaScoreDataLog) TableName() string {
	return "scoredatalog"
}

type BeatorajaScoreData struct {
	Sha256     string
	Mode       string
	Clear      int32
	Epg        int32
	Lpg        int32
	Egr        int32
	Lgr        int32
	Egd        int32
	Lgd        int32
	Ebd        int32
	Lbd        int32
	Epr        int32
	Lpr        int32
	Ems        int32
	Lms        int32
	Notes      int32
	Combo      int32
	Minbp      int32
	PlayCount  int32 `gorm:"column:playcount"`
	ClearCount int32 `gorm:"column:clearcount"`
	Trophy     string
	Ghost      string
	Option     int32
	Seed       int64
	Random     int32
	TimeStamp  int64 `gorm:"column:date"`
	State      int32
}

func (BeatorajaScoreData) TableName() string {
	return "score"
}

type BeatorajaSongInfo struct {
	Sha256       string
	N            int
	LN           int     `gorm:"column:ln"`
	S            int     `gorm:"column:s"`
	LS           int     `gorm:"column:ls"`
	Total        float64 `gorm:"column:total"`
	Density      float64 `gorm:"column:density"`
	PeakDensity  float64 `gorm:"column:peakdensity"`
	EndDensity   float64 `gorm:"column:enddensity"`
	MainBPM      float64 `gorm:"column:mainbpm"`
	Distribution string  `gorm:"column:distribution"`
	SpeedChange  string  `gorm:"column:speedchange"`
	LaneNotes    string  `gorm:"column:lanenotes"`
}

func (BeatorajaSongInfo) TableName() string {
	return "information"
}

type BeatorajaScanResult struct {
	SongData string
	SongInfo string
	Players  []BeatorajaPlayerFolder
}

type BeatorajaPlayerFolder struct {
	Name         string
	Path         string
	Score        string
	ScoreLog     string
	ScoreDataLog string
}
