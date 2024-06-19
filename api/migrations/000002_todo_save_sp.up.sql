CREATE OR REPLACE PROCEDURE public.todo_save(IN id integer, IN title text, IN completed boolean, id OUT integer)
 LANGUAGE sql
BEGIN ATOMIC
 INSERT INTO todos (title, completed)
   VALUES (todo_save.title, todo_save.completed)
   RETURNING id;
END