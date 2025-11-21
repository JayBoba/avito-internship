package dataflow

import "github.com/JayBoba/avito-internship/internal/entity"

// ----- TO DTO -----
func ToTeamDTO(team *entity.Team) TeamDTO {
	membersDTO := make([]TeamMemberDTO, len(team.Members))
	for i, member := range team.Members {
		membersDTO[i] = ToTeamMemberDTO(&member)
	}
	return TeamDTO{
		TeamName: team.TeamName,
		Members:  membersDTO,
	}
}

func ToTeamMemberDTO(teamMember *entity.TeamMember) TeamMemberDTO {
	return TeamMemberDTO{
		UserID:   teamMember.UserID,
		Username: teamMember.Username,
		IsActive: teamMember.IsActive,
	}
}

func ToUserDTO(user *entity.User) UserDTO {
	return UserDTO{
		UserID:   user.UserID,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}

func ToPullRequestDTO(pr *entity.PullRequest) PullRequestDTO {
	reviewersDTO := make([]UserDTO, len(pr.AssignedReviewers))
	for i, reviewer := range pr.AssignedReviewers {
		reviewersDTO[i] = ToUserDTO(&reviewer)
	}

	return PullRequestDTO{
		PRID:              pr.PRID,
		PRName:            pr.PRName,
		AuthorID:          pr.AuthorID,
		Status:            StatusToDTO(pr.Status),
		AssignedReviewers: reviewersDTO,
	}
}

func StatusToDTO(status entity.Status) Status {
	switch status {
	case entity.StatusOpen:
		return StatusOpen
	case entity.StatusMerged:
		return StatusMerged
	default:
		return StatusUnknown // ЭТОГО В АПИ ПОКА НЕТ НАДО ДУМАТЬ ЧТО ТУТ ДЕЛАТЬ
	}
}

//----- TO ENTITY -----

func ToTeam(teamDTO *TeamDTO) entity.Team {
	members := make([]entity.TeamMember, len(teamDTO.Members))
	for i, member := range teamDTO.Members {
		members[i] = ToTeamMember(&member)
	}
	return entity.Team{
		TeamName: teamDTO.TeamName,
		Members:  members,
	}
}

func ToTeamMember(teamMemberDTO *TeamMemberDTO) entity.TeamMember {
	return entity.TeamMember{
		UserID:   teamMemberDTO.UserID,
		Username: teamMemberDTO.Username,
		IsActive: teamMemberDTO.IsActive,
	}
}

func ToUser(userDTO *UserDTO) entity.User {
	return entity.User{
		UserID:   userDTO.UserID,
		Username: userDTO.Username,
		TeamName: userDTO.TeamName,
		IsActive: userDTO.IsActive,
	}
}

func ToPullRequest(prDTO *PullRequestDTO) entity.PullRequest {
	reviewers := make([]entity.User, len(prDTO.AssignedReviewers))
	for i, reviewer := range prDTO.AssignedReviewers {
		reviewers[i] = ToUser(&reviewer)
	}

	return entity.PullRequest{
		PRID:              prDTO.PRID,
		PRName:            prDTO.PRName,
		AuthorID:          prDTO.AuthorID,
		Status:            StatusToEntity(prDTO.Status),
		AssignedReviewers: reviewers,
	}
}

func StatusToEntity(statusDTO Status) entity.Status {
	switch statusDTO {
	case StatusOpen:
		return entity.StatusOpen
	case StatusMerged:
		return entity.StatusMerged
	default:
		return entity.StatusUnknown // ЭТОГО В ПОКА НЕТ. КОСТЫЛЬ. ПУСТАЯ СТРОКА
	}
}
