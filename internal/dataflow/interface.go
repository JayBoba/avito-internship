package dataflow

import "github.com/JayBoba/avito-internship/internal/entity"

//----- TEAM
type DefaultTeamMapper struct{}

type TeamMapper interface {
	ToTeamDTO(team *entity.Team) TeamDTO
	ToTeamMemberDTO(teamMember *entity.TeamMember) TeamMemberDTO
	ToTeam(teamDTO *TeamDTO) entity.Team
	ToTeamMember(teamMemberDTO *TeamMemberDTO) entity.TeamMember
}

func (m *DefaultTeamMapper) ToTeamDTO(team *entity.Team) TeamDTO {
	return ToTeamDTO(team)
}

func (m *DefaultTeamMapper) ToTeamMemberDTO(teamMember *entity.TeamMember) TeamMemberDTO {
	return ToTeamMemberDTO(teamMember)
}

func (m *DefaultTeamMapper) ToTeam(teamDTO *TeamDTO) entity.Team {
	return ToTeam(teamDTO)
}

func (m *DefaultTeamMapper) ToTeamMember(teamMemberDTO *TeamMemberDTO) entity.TeamMember {
	return ToTeamMember(teamMemberDTO)
}

//----- USER
type DefaultUserMapper struct{}

type UserMapper interface {
	ToUserDTO(user *entity.User) UserDTO
	ToUser(userDTO *UserDTO) entity.User
}

func (m *DefaultUserMapper) ToUserDTO(user *entity.User) UserDTO {
	return ToUserDTO(user)
}

func (m *DefaultUserMapper) ToUser(userDTO *UserDTO) entity.User {
	return ToUser(userDTO)
}

//----- PULL REQUESTS

type DefaultPRMapper struct{}

type PRMapper interface {
	ToPullRequestDTO(pr *entity.PullRequest) PullRequestDTO
	ToPullRequest(prDTO *PullRequestDTO) entity.PullRequest
}

func (m *DefaultPRMapper) ToPullRequestDTO(pr *entity.PullRequest) PullRequestDTO {
	return ToPullRequestDTO(pr)
}

func (m *DefaultPRMapper) ToPullRequest(prDTO *PullRequestDTO) entity.PullRequest {
	return ToPullRequest(prDTO)
}

//----- STATUS

type DefaultStatusMapper struct{}

type StatusMapper interface {
	StatusToDTO(status entity.Status) Status
	StatusToEntity(status Status) entity.Status
}

func (m *DefaultStatusMapper) StatusToDTO(status entity.Status) Status{
	return StatusToDTO(status)
}

func (m *DefaultStatusMapper) StatusToEntity(statusDTO Status) entity.Status{
	return StatusToEntity(statusDTO)
}