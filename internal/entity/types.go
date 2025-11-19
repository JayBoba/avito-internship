package entity

/*
    Team:
      type: object
      required: [ team_name, members]
      properties:
        team_name:
          type: string
        members:
          type: array
          items:
            $ref: '#/components/schemas/TeamMember'



    TeamMember:
      type: object
      required: [ user_id, username, is_active ]
      properties:
        user_id:
          type: string
        username:
          type: string
        is_active:
          type: boolean



    User:
      type: object
      required: [ user_id, username, team_name, is_active ]
      properties:
        user_id:
          type: string
        username:
          type: string
        team_name:
          type: string
        is_active:
          type: boolean



    PullRequest:
      type: object
      required: [ pull_request_id, pull_request_name, author_id, status, assigned_reviewers]
      properties:
        pull_request_id:
          type: string
        pull_request_name:
          type: string
        author_id:
          type: string
        status:
          type: string
          enum: [OPEN, MERGED]
        assigned_reviewers:
          type: array
          items:
            type: string
          description: user_id назначенных ревьюверов (0..2)
        createdAt:
          type: string
          format: date-time
          nullable: true
        mergedAt:
          type: string
          format: date-time
          nullable: true


		  
    PullRequestShort:
      type: object
      required: [ pull_request_id, pull_request_name, author_id, status]
      properties:
        pull_request_id:
          type: string
        pull_request_name:
          type: string
        author_id:
          type: string
        status:
          type: string
          enum: [OPEN, MERGED]
*/