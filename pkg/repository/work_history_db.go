package repository

import (
	"context"
	"database/sql"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/model/company"
	models "github.com/hiromaily/go-graphql-server/pkg/model/rdb"
	"github.com/hiromaily/go-graphql-server/pkg/model/workhistory"
)

type workHistoryDB struct {
	dbConn    *sql.DB
	tableName string
	logger    *zap.Logger
	company   company.Company
}

// NewWorkHistoryDBRepo returns WorkHistory interface
func NewWorkHistoryDBRepo(dbConn *sql.DB, logger *zap.Logger, company company.Company) workhistory.WorkHistory {
	return &workHistoryDB{
		dbConn:    dbConn,
		tableName: "t_user_work_history",
		logger:    logger,
		company:   company,
	}
}

// Fetch returns work history by userID
func (w *workHistoryDB) Fetch(userID string) ([]*workhistory.WorkHistoryType, error) {
	ctx := context.Background()

	var workHistories []*workhistory.WorkHistoryType
	err := models.TUserWorkHistories(
		qm.Select("wh.user_id, cp.name as company, wh.title, wh.description, wh.started_at, wh.started_at"),
		qm.From("t_user_work_history as wh"),
		qm.LeftOuterJoin("t_company as cp on wh.company_id = cp.id"),
		qm.Where("wh.user_id=?", userID),
	).Bind(ctx, w.dbConn, &workHistories)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUserWorkHistories().Bind() in Fetch()")
	}

	return workHistories, nil
}

// FetchAll returns all work history
func (w *workHistoryDB) FetchAll() ([]*workhistory.WorkHistoryType, error) {
	ctx := context.Background()

	var workHistories []*workhistory.WorkHistoryType
	err := models.TUserWorkHistories(
		qm.Select("wh.user_id, cp.name as company, wh.title, wh.description, wh.started_at, wh.started_at"),
		qm.From("t_user_work_history as wh"),
		qm.LeftOuterJoin("t_company as cp on wh.company_id = cp.id"),
	).Bind(ctx, w.dbConn, &workHistories)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUserWorkHistories().Bind() in Fetch()")
	}

	return workHistories, nil
}

func (w *workHistoryDB) Insert(wt *workhistory.WorkHistoryType) error {
	// get company
	companyType, err := w.company.FetchByName(wt.Company)
	if err != nil {
		return err
	}

	item := &models.TUserWorkHistory{
		UserID:      wt.UserID,
		CompanyID:   companyType.ID,
		Title:       wt.Title,
		Description: wt.Description,
	}
	if wt.TechIDs != nil {
		// convert []int to types.JSON(json.RawMessage)
		marshaled, err := jsoniter.Marshal(wt.TechIDs)
		if err != nil {
			return err
		}
		item.TechIds = marshaled
	}
	if wt.StartedAt != nil {
		item.CreatedAt = null.TimeFromPtr(wt.StartedAt)
	}
	if wt.EndedAt != nil {
		item.EndedAt = null.TimeFromPtr(wt.EndedAt)
	}

	ctx := context.Background()

	if err := item.Insert(ctx, w.dbConn, boil.Infer()); err != nil {
		return errors.Wrap(err, "failed to call user.Insert()")
	}
	// TODO: return latest ID
	return nil
}

func (w *workHistoryDB) Update(wt *workhistory.WorkHistoryType) error {
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	if wt.Company != "" {
		cp, err := w.company.FetchByName(wt.Company)
		if err != nil {
			return err
		}
		updCols[models.TUserWorkHistoryColumns.CompanyID] = cp.ID
	}
	if wt.Title != "" {
		updCols[models.TUserWorkHistoryColumns.Title] = wt.Title
	}
	if wt.Description != "" {
		updCols[models.TUserWorkHistoryColumns.Description] = wt.Description
	}
	if wt.TechIDs != nil {
		updCols[models.TUserWorkHistoryColumns.TechIds] = wt.TechIDs
	}
	if wt.StartedAt != nil {
		updCols[models.TUserWorkHistoryColumns.StartedAt] = wt.StartedAt
	}
	if wt.EndedAt != nil {
		updCols[models.TUserWorkHistoryColumns.EndedAt] = wt.EndedAt
	}
	updCols[models.TUserColumns.UpdatedAt] = null.TimeFrom(time.Now().UTC())

	_, err := models.TUsers(
		qm.Where("id=?", wt.ID),
	).UpdateAll(ctx, w.dbConn, updCols)

	return err
}

func (w *workHistoryDB) Delete(id string) error {
	ctx := context.Background()

	_, err := models.TUsers(
		qm.Where("t_user_work_history.id=?", id),
	).DeleteAll(ctx, w.dbConn)
	return err
}
