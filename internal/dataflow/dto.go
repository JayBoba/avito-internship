package dataflow

type TeamDTO struct {
	TeamName string
	Members  []TeamMemberDTO
}

type TeamMemberDTO struct {
	UserID   string
	Username string
	IsActive bool
}

type UserDTO struct {
	UserID   string
	Username string
	TeamName string
	IsActive bool
}

type PullRequestDTO struct {
	PRID     string
	PRName   string
	AuthorID string
	Status
	AssignedReviewers []UserDTO
}

type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusMerged Status = "MERGED"
)
