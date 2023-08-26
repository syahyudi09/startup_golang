package repo

import (
	"database/sql"
	"fmt"
	"startup/model"
	"startup/utils"
)

type CampaignRepo interface {
	FindAll() ([]*model.CampaignModel, error)
	FindByID(int) (*model.CampaignModel, error)
}

type campaignRepoImpl struct {
	db *sql.DB
}

func (cr *campaignRepoImpl) FindAll() ([]*model.CampaignModel, error) {
	// query untuk mendapatkan semua data campaign ini berada pada file utils.constant.go
	query := utils.FIND_CAMPAIGN_ALL

	// untuk melakukan pencarian/pengambilan data pada database
	rows, err := cr.db.Query(query)
	// untuk mengecek error
	if err != nil {
		return nil, fmt.Errorf("error on campaignRepoImpl.FindAll: %w", err)
	}
	/* 
	jika pengambilan data berhasil atau terjadi error
	maka akan mengakhiri pencarian atau pengambilan 
	data dengan defer rows.close dan menutup nya
	*/
	defer rows.Close()

	// deklarasi variabel yang bertindak sebagai penampung untuk data campaign
	var arryCampaign []*model.CampaignModel
	// 
	for rows.Next() {
		campaign := &model.CampaignModel{}
		campaignImages := model.CampaignImages{}
		if err := rows.Scan(
			&campaignImages.CampingID, &campaignImages.FileName, &campaignImages.IsPrimary,
			&campaign.ID, &campaign.UserID, &campaign.Name, &campaign.ShortDescription,
			&campaign.GoalAmount, &campaign.CurrentAmount, &campaign.Slug,
		); err != nil {
			return nil, fmt.Errorf("error scanning campaign row: %w", err)
		}
		campaign.CampaignImages = append(campaign.CampaignImages, campaignImages)
		arryCampaign = append(arryCampaign, campaign)
		
	}
	return arryCampaign, nil
}

func (cr *campaignRepoImpl) FindByID(userID int) (*model.CampaignModel, error) {
	query := utils.FIND_CAMPAIGN_BY_ID

	rows, err := cr.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("err on campaignRepoImpl.FindByID %w", err)
	}
	defer rows.Close()

	campaign := &model.CampaignModel{}
	for rows.Next() {
		campaignImages := model.CampaignImages{}
		if err := rows.Scan(
			&campaignImages.CampingID, &campaignImages.FileName, &campaignImages.IsPrimary, 
			&campaign.ID, &campaign.UserID, &campaign.Name, &campaign.ShortDescription,
			&campaign.GoalAmount, &campaign.CurrentAmount, &campaign.Slug,
		); err != nil {
			return nil, fmt.Errorf("err on campaignRepoImpl.FindByID %w", err)
		}
		campaign.CampaignImages = append(campaign.CampaignImages, campaignImages)
	}

	if campaign.UserID == 0{
		return nil, nil
	}

	return campaign, nil
}

func NewCampaignRepo(db *sql.DB) CampaignRepo {
	return &campaignRepoImpl{
		db: db,
	}
}