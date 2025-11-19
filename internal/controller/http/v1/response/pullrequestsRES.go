package response

/* post open
      responses:
        '201':
          description: PR создан
          content:
            application/json:
              schema:
                type: object
                properties:
                  pr:
                    $ref: '#/components/schemas/PullRequest'
              example:
                pr:
                  pull_request_id: pr-1001
                  pull_request_name: Add search
                  author_id: u1
                  status: OPEN
                  assigned_reviewers: [u2, u3]
        '404':
          description: Автор/команда не найдены
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
        '409':
          description: PR уже существует
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
              example:
                error: { code: PR_EXISTS, message: PR id already exists }
*/

/* post merge
      responses:
        '200':
          description: PR в состоянии MERGED
          content:
            application/json:
              schema:
                type: object
                properties:
                  pr:
                    $ref: '#/components/schemas/PullRequest'
              example:
                pr:
                  pull_request_id: pr-1001
                  pull_request_name: Add search
                  author_id: u1
                  status: MERGED
                  assigned_reviewers: [u2, u3]
                  mergedAt: 2025-10-24T12:34:56Z
        '404':
          description: PR не найден
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
*/

/* post reassign
      responses:
        '200':
          description: Переназначение выполнено
          content:
            application/json:
              schema:
                type: object
                required: [pr, replaced_by]
                properties:
                  pr:
                    $ref: '#/components/schemas/PullRequest'
                  replaced_by:
                    type: string
                    description: user_id нового ревьювера
              example:
                pr:
                  pull_request_id: pr-1001
                  pull_request_name: Add search
                  author_id: u1
                  status: OPEN
                  assigned_reviewers: [u3, u5]
                replaced_by: u5


		'404':
          description: PR или пользователь не найден
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
        
		
		'409':
          description: Нарушение доменных правил переназначения
          content:
            application/json:
              schema: { $ref: '#/components/schemas/ErrorResponse' }
              examples:
                merged:
                  summary: Нельзя менять после MERGED
                  value:
                    error: { code: PR_MERGED, message: cannot reassign on merged PR }
                notAssigned:
                  summary: Пользователь не был назначен ревьювером
                  value:
                    error: { code: NOT_ASSIGNED, message: reviewer is not assigned to this PR }
                noCandidate:
                  summary: Нет доступных кандидатов
                  value:
                    error: { code: NO_CANDIDATE, message: no active replacement candidate in team }

*/