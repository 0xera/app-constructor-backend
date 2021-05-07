create table if not exists project
(
    id   serial primary key not null,
    name varchar unique     not null,
    app  varchar default '{}'
);


create table if not exists user_data
(
    id            varchar primary key not null,
    widgets_count int default 0,
    email         varchar             not null
);

create table if not exists user_projects
(
    user_id    varchar references user_data (id) on delete cascade,
    project_id serial references project (id) on delete cascade
);

create unique index user_index on user_data (id, email);