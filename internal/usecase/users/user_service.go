package users

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type UserService interface {
	All (ctx context.Context) ([]User, error)
	Search (ctx context.Context, filter *UserFilter) (*SearchResult, error)
}

func NewUserService(db *sql.DB) UserService {
	return &userService{db: db}
}

type userService struct {
	db	*sql.DB
}

func (s *userService)  All(ctx context.Context) ([]User, error){
	query := "Select * from users"
	res, err := s.db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	userRes, err := userResult(res)
	return userRes, nil
}

func (s *userService)  Search(ctx context.Context, filter *UserFilter) (*SearchResult, error){
	totalUser, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	query, params := BuildSearchQuery(filter)
	res, err := s.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}

	userRes, err := userResult(res)
	var searchRes SearchResult
	searchRes.List = userRes
	searchRes.Total = len(totalUser)
	searchRes.PageSize = filter.PageSize
	searchRes.PageIndex = filter.PageIndex
	return &searchRes, nil
}

func userResult(query *sql.Rows) ([]User, error) {
	var res []User
	user := User{}
	for query.Next() {
		err := query.Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}

func BuildSearchQuery(filter *UserFilter) (string, []interface{}) {
	query := "select * from users"
	var condition []string
	var params []interface{}
	i := 1

	if len(filter.Id) > 0 {
		params = append(params, filter.Id)
		condition = append(condition, fmt.Sprintf(`id = $%d`, i))
		i++
	}
	if len(filter.Email) > 0 {
		q := "%" + filter.Email + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`email ilike $%d`, i))
		i++
	}
	if len(filter.Username) > 0 {
		q := "%" + filter.Username + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`username ilike $%d`, i))
		i++
	}
	if len(filter.Phone) > 0 {
		q := "%" + filter.Phone + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`phone ilike $%d`, i))
		i++
	}

	if len(condition) > 0 {
		cond := strings.Join(condition, " and ")
		query += fmt.Sprintf(` where %s order by id`, cond)
	}

	if filter.PageSize > 0 {
		query = query + fmt.Sprintf(` limit %d`, filter.PageSize)
		if filter.PageIndex > 0 {
			pageIndex := (filter.PageIndex - 1) * filter.PageSize
			query = query + fmt.Sprintf(` offset %d`, pageIndex)
		}
	}

	return query, params
}