package workhistory

import "time"

// WorkHistory for fetching data interface
// - implementation is in repository
type WorkHistory interface {
	Fetch(userID string) ([]*WorkHistoryType, error)
	FetchAll() ([]*WorkHistoryType, error)
	Insert(wt *WorkHistoryType) error
	Update(wt *WorkHistoryType) error
	Delete(userID int) error
}

// WorkHistoryType is type of work history
type WorkHistoryType struct {
	ID          int        `json:"id" boil:"id"`
	UserID      int        `json:"user_id" boil:"user_id"`
	Company     string     `json:"company" boil:"company"`
	Title       string     `json:"title" boil:"title"`
	Description string     `json:"description" boil:"description"`
	TechIDs     []int      `json:"tech_ids" boil:"tech_ids"`
	StartedAt   *time.Time `json:"started_at" boil:"started_at"`
	EndedAt     *time.Time `json:"ended_at" boil:"ended_at"`
}
