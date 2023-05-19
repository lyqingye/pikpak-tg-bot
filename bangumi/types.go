package bangumi

import "errors"

const (
	Resolution1080p   = "1080P"
	Resolution720p    = "720P"
	ResolutionUnknown = "unknown"

	SubtitleChs     = "CHS"
	SubtitleCht     = "CHT"
	SubtitleUnknown = "unknown"
)

type Bangumi struct {
	Title     string
	EPCount   uint
	Season    uint
	TmDBId    int64
	SubjectId int64
	Episodes  []Episode
}

type Episode struct {
	BangumiTitle string
	SubjectId    int64 // BangumiTV SubjectId
	EpisodeTitle string
	Subgroup     string
	Season       uint
	EPNumber     uint
	Magnet       string
	TorrentHash  string
	Torrent      []byte
	Date         string
	FileSize     uint64
	Lang         []string
	Resolution   string
	Read         bool // mark as read
}

func (e *Episode) Validate() error {
	if e.EPNumber <= 0 {
		return errors.New("invalid ep number")
	}
	if e.Magnet == "" && len(e.Torrent) == 0 {
		return errors.New("empty download resource")
	}
	if e.FileSize == 0 {
		return errors.New("invalid filesize")
	}
	if e.BangumiTitle == "" {
		return errors.New("empty bangumi title")
	}
	return nil
}
