package response

/* post
      responses:
        '201':
          description: Команда создана
          content:
            application/json:
              schema:
                type: object
                properties:
                  team:
                    $ref: '#/components/schemas/Team'
              example:
                team:
                  team_name: backend
                  members:
                    - user_id: u1
                      username: Alice
                      is_active: true
                    - user_id: u2
                      username: Bob
                      is_active: true


'400':
          description: Команда уже существует
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
              example:
                error:
                  code: TEAM_EXISTS
                  message: team_name already exists
*/

/* get
      responses:
        '200':
          description: Объект команды
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Team'
              example:
                team_name: backend
                members:
                  - user_id: u1
                    username: Alice
                    is_active: true
                  - user_id: u2
                    username: Bob
                    is_active: true
        '404':
          description: Команда не найдена
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
*/