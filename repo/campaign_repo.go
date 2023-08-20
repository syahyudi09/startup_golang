package repo

import "database/sql"

type CampaignRepo interface {
}

type campaignRepoImpl struct {
	db *sql.DB
}

func NewCampaignRepo(db *sql.DB) CampaignRepo {
	return &campaignRepoImpl{
		db: db,
	}
}