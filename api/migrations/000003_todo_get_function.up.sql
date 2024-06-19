CREATE OR REPLACE FUNCTION public.todo_get(_id integer)
 RETURNS SETOF todos
 LANGUAGE sql
AS $function$
 
 select * from todos where id=_id;
 
$function$
;