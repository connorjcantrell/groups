package form

import (
	"encoding/json"
	"time"
)

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
	BookID    int `json:"book_id"`
	ChapterID int `json:"chapter_id"`
	SectionID int `json:"section_id"`
	// Client must send rfc3339/ISO 8601 string
	// json.Unmarshal() will convert to time.Time
	StartTime   time.Time `json:"start_time"`
	Description string    `json:"description"`
}

func UnmarshalEventForm(data []byte) (EventForm, error) {
	var f EventForm
	err := json.Unmarshal(data, &f)
	return f, err
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
	if f.StartTime.IsZero() {
		return false, NoStartTimeErr{}
	}
	return true, nil
}
