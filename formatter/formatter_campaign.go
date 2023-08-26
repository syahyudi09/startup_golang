package formatter

import (
	"startup/model"
)

type CampaignFormatter struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Name            string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl        string `json:"image_url"`
	GoalAmount      int    `json:"goal_amount"`
	CurrentAmount   int    `json:"current_amount"`
	Slug			string `json:"slug"`
}

func FormatCampaign(campaign *model.CampaignModel) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:              campaign.ID,
		UserID:          campaign.UserID,
		Name:            campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:      campaign.GoalAmount,
		CurrentAmount:   campaign.CurrentAmount,
		Slug: campaign.Slug,
		ImageUrl: "",
	}
	
	if campaign.CampaignImages != nil && len(campaign.CampaignImages) > 0 {
		// Mengiterasi melalui CampaignImages yang terkait
		for _, img := range campaign.CampaignImages {
			// Jika is_primary adalah 1, gunakan FileName sebagai ImageUrl
			if img.IsPrimary == 1 {
				campaignFormatter.ImageUrl = img.FileName
			}
			 // Keluar dari loop karena kita sudah menemukan gambar utama
			if img.IsPrimary == 0 {
				campaignFormatter.ImageUrl = "" // Jika is_primary adalah 0, set ImageUrl kosong
			}
		}
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []*model.CampaignModel) []CampaignFormatter {
	formattedCampaigns := make([]CampaignFormatter, 0)

	for _, campaign := range campaigns {
		formatted := FormatCampaign(campaign)
		formattedCampaigns = append(formattedCampaigns, formatted)
	}

	return formattedCampaigns
}
