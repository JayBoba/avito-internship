package response

/* post
      responses:
        '200':
          description: Обновлённый пользователь
          content:
            application/json:
              schema:
                type: object
                properties:
                  user:
                    $ref: '#/components/schemas/User'
              example:
                user:
                  user_id: u2
                  username: Bob
                  team_name: backend
                  is_active: false
        '404':
          description: Пользователь не найден
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
*/

/* get
      responses:
        '200':
          description: Список PR'ов пользователя
          content:
            application/json:
              schema:
                type: object
                required: [ user_id, pull_requests ]
                properties:
                  user_id:
                    type: string
                  pull_requests:
                    type: array
                    items:
                      $ref: '#/components/schemas/PullRequestShort'
              example:
                user_id: u2
                pull_requests:
                  - pull_request_id: pr-1001
                    pull_request_name: Add search
                    author_id: u1
                    status: OPEN
*/