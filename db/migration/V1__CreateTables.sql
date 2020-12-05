create table program (program_id serial, program_name varchar(50));

create user webserviceuser with password 'webserviceuserpwd';

grant select, update, insert, delete on public.program to webserviceuser;

grant select, update on public.program_program_id_seq to webserviceuser;



