package sqlstore

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v6/model"
)

var escapeLikeSearchChar = []string{
	"%",
	"_",
}

func sanitizeSearchTerm(term string) string {
	const escapeChar = "\\"

	term = strings.Replace(term, escapeChar, "", -1)

	for _, c := range escapeLikeSearchChar {
		term = strings.Replace(term, c, escapeChar+c, -1)
	}

	return term
}

func (ss SQLStore) SearchGroupsByPrefix(prefix string) ([]*model.Group, error) {
	sanitizedPrefix := strings.ToLower(sanitizeSearchTerm(prefix))
	query := ss.replicaBuilder.
		Select("Id", "DisplayName", "DeleteAt").
		From("UserGroups").
		Where(sq.Like{"LOWER(DisplayName)": sanitizedPrefix + "%"}).
		Where(sq.Eq{"DeleteAt": 0}).
		OrderBy("DisplayName").
		Limit(10)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build SQL query for searching groups")
	}

	var groups []*model.Group
	if err := ss.replica.Select(&groups, sql, args...); err != nil {
		return nil, errors.Wrap(err, "failed to search groups")
	}

	return groups, nil
}
