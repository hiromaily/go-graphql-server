package workhistory

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// WorkHistory for fetching data interface
// - implementation is in repository
type WorkHistory interface {
	Fetch(userID string) ([]*WorkHistoryType, error)
	FetchAll() ([]*WorkHistoryType, error)
	Insert(wt *WorkHistoryType) error
	Update(wt *WorkHistoryType) error
	Delete(id string) error
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

// WorkHistoryFieldResolver for resolver of schema interface
type WorkHistoryFieldResolver interface {
	GetByUserID(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Create(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
	Delete(p graphql.ResolveParams) (interface{}, error)
}

type workHistoryFieldResolver struct {
	logger          *zap.Logger
	workHistoryRepo WorkHistory
}

// NewWorkHistoryFieldResolve returns WorkHistoryFieldResolver interface
func NewWorkHistoryFieldResolve(
	logger *zap.Logger,
	workHistoryRepo WorkHistory,
) WorkHistoryFieldResolver {
	return &workHistoryFieldResolver{
		logger:          logger,
		workHistoryRepo: workHistoryRepo,
	}
}

// GetByUserID gets work history by UserID
func (w *workHistoryFieldResolver) GetByUserID(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOK := p.Args["user_id"].(string)
	if isOK {
		return w.workHistoryRepo.Fetch(idQuery)
	}
	return nil, errors.New("not found")
}

// List returns all work history
func (w *workHistoryFieldResolver) List(_ graphql.ResolveParams) (interface{}, error) {
	return w.workHistoryRepo.FetchAll()
}

// Create creates new work history by parameters
func (w *workHistoryFieldResolver) Create(p graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	newWorkHistory := &WorkHistoryType{
		ID:          rand.Intn(100000), // TODO: get maximum ID from list
		UserID:      p.Args["user_id"].(int),
		Company:     p.Args["company"].(string),
		Title:       p.Args["title"].(string),
		Description: p.Args["description"].(string),
		TechIDs:     []int{},
		StartedAt:   p.Args["started_at"].(*time.Time),
	}
	if v, ok := p.Args["started_at"].(*time.Time); ok {
		newWorkHistory.EndedAt = v
	}

	// insert to repository
	err := w.workHistoryRepo.Insert(newWorkHistory)
	if err != nil {
		return nil, err
	}
	return newWorkHistory, nil
}

func (w *workHistoryFieldResolver) Update(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	updated := WorkHistoryType{
		ID: intID,
	}

	//if userID, ok := p.Args["user_id"].(int); ok {
	//	updated.UserID = userID
	//}
	//TODO: may be better to update by company ID
	if company, ok := p.Args["company"].(string); ok {
		updated.Company = company
	}
	if company, ok := p.Args["company"].(string); ok {
		updated.Company = company
	}
	if title, ok := p.Args["title"].(string); ok {
		updated.Title = title
	}
	if description, ok := p.Args["description"].(string); ok {
		updated.Description = description
	}
	if startedAt, ok := p.Args["started_at"].(*time.Time); ok {
		updated.StartedAt = startedAt
	}
	if startedAt, ok := p.Args["ended_at"].(*time.Time); ok {
		updated.StartedAt = startedAt
	}
	if err := w.workHistoryRepo.Update(&updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (w *workHistoryFieldResolver) Delete(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	deleted, err := w.workHistoryRepo.Fetch(id)
	if err != nil {
		return nil, err
	}
	w.workHistoryRepo.Delete(id)

	return deleted, nil
}
