package pgbuilder

import (
	"github.com/kunlun-qilian/sqlx/v3"
	"github.com/kunlun-qilian/sqlx/v3/builder"
)

type DataList interface {
	sqlx.ScanIterator
	ConditionBuilder
	DoList(db sqlx.DBExecutor, pager *Pager, additions ...builder.Addition) error
}

func BatchDoList(db sqlx.DBExecutor, scanners ...DataList) (err error) {
	if len(scanners) == 0 {
		return nil
	}

	for i := range scanners {
		scanner := scanners[i]

		cond := scanner.ToCondition(db)

		if !builder.IsNilExpr(cond) {
			if err := scanner.DoList(db, nil); err != nil {
				return err
			}
		}
	}

	return nil
}
