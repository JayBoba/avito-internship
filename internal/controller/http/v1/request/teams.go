package request

/* 
package request

/* 
  /team/add:
    post:
      tags: [Teams]
      summary: Создать команду с участниками (создаёт/обновляет пользователей)
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Team'
            example:
              team_name: payments
              members:
                - user_id: u1
                  username: Alice
                  is_active: true
                - user_id: u2
                  username: Bob
                  is_active: true
*/

/*
  /team/get:
    get:
      tags: [Teams]
      summary: Получить команду с участниками
      parameters:
        - $ref: '#/components/parameters/TeamNameQuery'
*/