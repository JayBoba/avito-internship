package request

/*
  /users/setIsActive:
    post:
      tags: [Users]
      summary: Установить флаг активности пользователя
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required: [ user_id, is_active ]
              properties:
                user_id:
                  type: string
                is_active:
                  type: boolean
            example:
              user_id: u2
              is_active: false
*/

/*
  /users/getReview:
    get:
      tags: [Users]
      summary: Получить PR'ы, где пользователь назначен ревьювером
      parameters:
        - $ref: '#/components/parameters/UserIdQuery'
*/