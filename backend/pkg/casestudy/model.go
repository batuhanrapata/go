package casestudy

import "time"

type CaseStudy struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURI    string    `json:"imageuri"`
	CreatedDate time.Time `json:"createddate" gorm:"autoCreateTime"`
}
