package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/country"
	models "github.com/hiromaily/go-graphql-server/pkg/model/rdb"
	"github.com/hiromaily/go-graphql-server/pkg/user"
)

type userDB struct {
	dbConn    *sql.DB
	tableName string
	logger    *zap.Logger
	country   country.Country
}

// NewUserDBRepo returns User interface
func NewUserDBRepo(dbConn *sql.DB, logger *zap.Logger, country country.Country) user.User {
	return &userDB{
		dbConn:    dbConn,
		tableName: "t_user",
		logger:    logger,
		country:   country,
	}
}

// Fetch returns user by id
func (u *userDB) Fetch(id string) (*user.UserType, error) {
	ctx := context.Background()

	var user user.UserType
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	u.logger.Debug("userdb.fetch 1")
	err = models.TUsers(
		qm.Select("t_user.id, t_user.name, t_user.age, cty.name as country"),
		qm.LeftOuterJoin("m_country as cty on t_user.country_id = cty.id"),
		qm.Where("t_user.id=?", intID),
	).Bind(ctx, u.dbConn, &user)
	u.logger.Debug("userdb.fetch 2")
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUsers().Bind() in Fetch()")
	}
	u.logger.Debug("userdb.fetch", zap.Any("user", user))

	return &user, nil
}

// FetchAll returns all users
func (u *userDB) FetchAll() ([]*user.UserType, error) {
	ctx := context.Background()

	var users []*user.UserType
	// sql := "SELECT id FROM t_users WHERE delete_flg=?"
	err := models.TUsers(
		qm.Select("t_user.id, t_user.name, t_user.age, cty.name as country"),
		qm.LeftOuterJoin("m_country as cty on t_user.country_id = cty.id"),
	).Bind(ctx, u.dbConn, &users)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUsers().Bind() in FetchAll()")
	}
	return users, nil
}

func (u *userDB) Insert(ut *user.UserType) error {
	// get country
	countryType, err := u.country.FetchByName(ut.Country)
	if err != nil {
		return err
	}

	item := &models.TUser{
		Name:      ut.Name,
		Age:       uint8(ut.Age),
		CountryID: uint8(countryType.ID),
	}

	ctx := context.Background()

	if err := item.Insert(ctx, u.dbConn, boil.Infer()); err != nil {
		return errors.Wrap(err, "failed to call user.Insert()")
	}
	// TODO: return latest ID
	return nil
}

// TODO: implement
func (u *userDB) Update(ut *user.UserType) error {
	return nil
}

// TODO: implement
func (u *userDB) Delete(id string) {
}
