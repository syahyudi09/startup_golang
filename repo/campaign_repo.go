package repo

import (
	"database/sql"
	"fmt"
	"startup/model"
)

type CampaignRepo interface {
	FindAll() ([]*model.CampaignModel, error)
	FindByID(int) (*model.CampaignModel, error)
}

type campaignRepoImpl struct {
	db *sql.DB
}

func (cr *campaignRepoImpl) FindAll() ([]*model.CampaignModel, error) {
	query := "SELECT id, user_id, name, short_description, description, goal_amount, current_amount, perks, backer_count, slug, created_at, updated_at FROM campaign"

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error on campaignRepoImpl.FindAll: %w", err)
	}
	defer rows.Close()

	var arryCampaign []*model.CampaignModel
	for rows.Next() {
		campaign := new(model.CampaignModel)
		if err := rows.Scan(
			&campaign.ID,
			&campaign.UserID,
			&campaign.Name,
			&campaign.ShortDescription,
			&campaign.Description,
			&campaign.GoalAmount,
			&campaign.CurrentAmount,
			&campaign.Perks,
			&campaign.BackerCount,
			&campaign.Slug,
			&campaign.CreateAt,
			&campaign.UpdateAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning campaign row: %w", err)
		}
		arryCampaign = append(arryCampaign, campaign)
	}
	return arryCampaign, nil
}

func (cr *campaignRepoImpl) FindByID(campaignID int) (*model.CampaignModel, error) {
	query := `SELECT ci.campaign_id, ci.is_primary, c.id, c.user_id, c.name, c.short_description, c.description, c.goal_amount, c.current_amount, c.perks, c.backer_count, c.slug, c.created_at, c.updated_at FROM campaign AS c JOIN campaign_images AS ci ON c.id = ci.campaign_id WHERE c.id = $1`
	
	campaign := &model.CampaignModel{}

	rows, err := cr.db.Query(query, campaignID)
	if err != nil {
		return nil, fmt.Errorf("err on campaignRepoImpl.FindByID %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		campaignImages := model.CampaignImages{}
		err := rows.Scan(
			&campaignImages.CampingID, &campaignImages.IsPrimary,
			&campaign.ID, &campaign.UserID, &campaign.Name, &campaign.ShortDescription, &campaign.Description, &campaign.GoalAmount, &campaign.CurrentAmount,
			&campaign.Perks, &campaign.BackerCount, &campaign.Slug, &campaign.CreateAt, &campaign.UpdateAt,
			// &campaignImages.FileName, &campaignImages.IsPrimary, &campaignImages.CreateAt, &campaignImages.UpdateAt,
		)
		if err != nil {
			return nil, fmt.Errorf("err on campaignRepoImpl.FindByID %w", err)
		}
		campaign.CampaignImages = append(campaign.CampaignImages, campaignImages)
	}
	return campaign, nil
}

func NewCampaignRepo(db *sql.DB)     CampaignRepo {
	return &campaignRepoImpl{
		db: db,
	}
}