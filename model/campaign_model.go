package model

import "time"

type CampaignModel struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreateAt         time.Time
	UpdateAt		 time.Time
	CampaignImages	[]CampaignImages
}

type CampaignImages struct{
	ID int
	CampingID int
	FileName string
	IsPrimary int
	CreateAt time.Time
	UpdateAt time.Time
}