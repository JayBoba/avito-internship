package request

/*
/pullRequest/create:
    post:
      tags: [PullRequests]
      summary: Создать PR и автоматически назначить до 2 ревьюверов из команды автора
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [ pull_request_id, pull_request_name, author_id ]
              properties:
                pull_request_id: { type: string }
                pull_request_name: { type: string }
                author_id: { type: string }
            example:
              pull_request_id: pr-1001
              pull_request_name: Add search
              author_id: u1
*/

/*
/pullRequest/merge:
    post:
      tags: [PullRequests]
      summary: Пометить PR как MERGED (идемпотентная операция)
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [ pull_request_id ]
              properties:
                pull_request_id: { type: string }
            example:
              pull_request_id: pr-1001
*/

/*
  /pullRequest/reassign:
    post:
      tags: [PullRequests]
      summary: Переназначить конкретного ревьювера на другого из его команды
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [ pull_request_id, old_user_id ]
              properties:
                pull_request_id: { type: string }
                old_user_id: { type: string }
            example:
              pull_request_id: pr-1001
              old_reviewer_id: u2
*/