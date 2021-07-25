package form

type NoBookErr struct{}

func (e NoBookErr) Error() string {
	return "no book selected"
}

type NoChapterErr struct{}

func (e NoChapterErr) Error() string {
	return "no chapter selected"
}

type NoSectionErr struct{}

func (e NoSectionErr) Error() string {
	return "no section selected"
}

type NoStartTimeErr struct{}

func (e NoStartTimeErr) Error() string {
	return "no start time given"
}

type NoDurationErr struct{}

func (e NoDurationErr) Error() string {
	return "no duration given"
}

type EventForm struct {
	BookID      int
	ChapterID   int
	SectionID   int
	StartTime   string // Convert formatted string to time.Time
	Description string
}

func (f EventForm) Validate() (bool, error) {
	if f.BookID == 0 {
		return false, NoBookErr{}
	}
	if f.ChapterID == 0 {
		return false, NoChapterErr{}
	}
	if f.SectionID == 0 {
		return false, NoSectionErr{}
	}
	if f.StartTime == "" {
		return false, NoStartTimeErr{}
	}
	return true, nil
}
