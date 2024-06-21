CREATE OR REPLACE PROCEDURE public.todo_delete(INOUT _id integer)
 LANGUAGE sql
BEGIN ATOMIC
 WITH deleted AS (
          DELETE FROM todos
           WHERE (todos.id = todo_delete._id)
           RETURNING todos.id,
             todos.title,
             todos.completed,
             todos.created_at,
             todos.updated_at
         )
  SELECT count(*) AS recordsaffected
    FROM deleted;
END
;