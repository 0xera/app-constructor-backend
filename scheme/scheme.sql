create table if not exists project
(
    _id  serial unique not null,
    id   int           not null,
    name varchar       not null,
    app  varchar default '{}',
    PRIMARY KEY (_id)
    );


create table if not exists user_data
(
    id            int primary key not null,
    widgets_count int default 0,
    email         varchar         not null
);

create table if not exists user_projects
(
    user_id    serial references user_data (id) on delete cascade,
    project_id serial references project (_id) on delete cascade
    );

create unique index user_index on user_data (id, email);


select *
from user_data
where id = ?;


-- projects
select *
from project
         inner join user_projects up on project._id = up.project_id
where up.user_id = ?;


-- projects
select *
from project
         inner join user_projects up on project._id = up.project_id
where up.user_id = ?
  and project.id = ?
    LIMIT 1;

select _id
from project
    inner join user_projects up on project._id = up.project_id
where up.user_id = ?
  and id = ?;


-- insert
insert into user_data(id, email)
values (?, ?);

insert into project(id, name)
values (?, ?)
    returning _id;


update project
set app = ?
where _id =?;

update user_data
set widgets_count = ?
where id =?;

delete
from project
where _id = ?



