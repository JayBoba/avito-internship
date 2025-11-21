package dataflow

import "github.com/JayBoba/avito-internship/internal/entity"

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
