package mysqlerrhandler

import (
	"database/sql"
	interror "market-starter/internal/error"
)

func NotFound(err error) bool {
	if err == nil {
		return false
	}

	if err.Error() == sql.ErrNoRows.Error() {
		return true
	}

	interror.Check(err)

	return false
}
