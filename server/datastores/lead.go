package datastores

import "github.com/jmoiron/sqlx"

const createLeadStm = `
	INSERT INTO leads(id, mail, hashCode, invited)
	VALUES(nextval('seq_leads'), $1, md5(to_char(now(), 'YYYYMMDDHHMISSMS')), $2)
`
const getLeadStm = `
select hashCode from leads where mail = $1
`
const countInvitesQuery = `
	SELECT count(1)
	FROM leads
	WHERE invited = $1
`

// Lead datastore
type Lead interface {
	Create(mail string, invited string) error
	CountByInvites(hashCode string) (int64, error)
	GetHashByMail(mail string) (string, error)
}

type lead struct {
	db *sqlx.DB
}

// NewLead datastore
func NewLead(db *sqlx.DB) Lead {
	return lead{db: db}
}

// Create a new lead in database.
func (ds lead) Create(mail string, invited string) error {
	_, err := ds.db.Exec(createLeadStm, mail, invited)
	return err
}

// Count leads by hashCode
func (ds lead) CountByInvites(hashCode string) (int64, error) {
	var count int64
	return count, ds.db.Get(&count, countInvitesQuery, hashCode)
}

// Get hashCode by mail
func (ds lead) GetHashByMail(mail string) (string, error) {
	var hashCode string
	return hashCode, ds.db.Get(&hashCode, getLeadStm, mail)
}
