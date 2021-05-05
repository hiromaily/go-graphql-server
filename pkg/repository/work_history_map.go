package repository

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/files"
	"github.com/hiromaily/go-graphql-server/pkg/model/workhistory"
)

type workHistoryMap struct {
	repo map[string][]*workhistory.WorkHistoryType // key is userID
	list []*workhistory.WorkHistoryType
}

// NewWorkHistoryMapRepo returns User interface
func NewWorkHistoryMapRepo() (workhistory.WorkHistory, error) {
	var data map[string][]*workhistory.WorkHistoryType
	err := files.ImportJSONFile("./assets/work_history.json", &data)
	if err != nil {
		return nil, err
	}
	return &workHistoryMap{
		repo: data,
	}, nil
}

func (w *workHistoryMap) updateList() {
	utList := make([]*workhistory.WorkHistoryType, 0, len(w.repo))
	for _, val := range w.repo {
		for _, v := range val {
			v := v
			utList = append(utList, v)
		}
	}
	w.list = utList
}

// Fetch returns user by id
func (w *workHistoryMap) Fetch(id string) (*workhistory.WorkHistoryType, error) {
	for _, vals := range w.repo {
		for _, v := range vals {
			if strconv.Itoa(v.ID) == id {
				v := v
				return v, nil
			}
		}
	}
	return nil, errors.New("work history is not found")
}

// FetchByUserID returns user by id
func (w *workHistoryMap) FetchByUserID(userID string) ([]*workhistory.WorkHistoryType, error) {
	if v, ok := w.repo[userID]; ok {
		return v, nil
	}
	return nil, errors.New("work history is not found")
}

// FetchAll returns all users
func (w *workHistoryMap) FetchAll() ([]*workhistory.WorkHistoryType, error) {
	if len(w.list) == 0 {
		w.updateList()
	}
	return w.list, nil
}

func (w *workHistoryMap) Insert(wt *workhistory.WorkHistoryType) error {
	userID := strconv.Itoa(wt.UserID)
	if _, ok := w.repo[userID]; ok {
		return errors.Errorf("id[%d] is already existing", wt.ID)
	}
	w.repo[userID] = append(w.repo[userID], wt)
	w.list = append(w.list, wt)

	return nil
}

func (w *workHistoryMap) Update(wt *workhistory.WorkHistoryType) error {
	updated, err := w.Fetch(strconv.Itoa(wt.ID))
	if err != nil {
		return errors.Wrapf(err, "id[%d] is not found", wt.ID)
	}

	if wt.Company != "" {
		wt.Company = updated.Company
	}
	if wt.Title != "" {
		wt.Company = updated.Company
	}
	if wt.Description != "" {
		wt.Company = updated.Description
	}
	if wt.TechIDs != nil {
		wt.TechIDs = updated.TechIDs
	}
	if wt.StartedAt != nil {
		wt.StartedAt = updated.StartedAt
	}
	if wt.EndedAt != nil {
		wt.EndedAt = updated.EndedAt
	}
	for key, vals := range w.repo {
		for idx, v := range vals {
			if v.ID == wt.ID {
				w.repo[key][idx] = wt
			}
		}
	}
	w.updateList()

	return nil
}

func (w *workHistoryMap) Delete(id string) error {
	// delelte
	for key, vals := range w.repo {
		for idx, v := range vals {
			if strconv.Itoa(v.ID) == id {
				// delete
				w.repo[key] = removeIndex(w.repo[key], idx)
			}
		}
	}
	w.updateList()
	return nil
}

func removeIndex(s []*workhistory.WorkHistoryType, index int) []*workhistory.WorkHistoryType {
	return append(s[:index], s[index+1:]...)
}
