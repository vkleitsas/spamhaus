package domain

type URLEntry struct {
	WebsiteURL      string `json:"WebsiteURL" bson:"WebsiteURL"`
	SubmissionCount int32  `json:"SubmissionCount" bson:"SubmissionCount"`
	Data            []byte `json:"Data" bson:"Data"`
}

type URLRequestInterface interface {
	GetBySize() ([]URLEntry, error)
	GetByDate() ([]URLEntry, error)
	GetMostSubmitted() ([]URLEntry, error)
	Update(d URLEntry) (*URLEntry, error)
	Delete(d URLEntry) error
}
