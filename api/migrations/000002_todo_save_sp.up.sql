CREATE OR REPLACE PROCEDURE todo_save(INOUT id integer, IN title text, IN completed boolean)
 LANGUAGE sql
BEGIN ATOMIC
 INSERT INTO todos (title, completed)
   VALUES (todo_save.title, todo_save.completed)
   RETURNING todos.id;
END
;