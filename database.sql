create table segments
(
    id       serial
        primary key,
    seg_name varchar(255)
);

alter table segments
    owner to sivovgg;

create table user_segment
(
    user_id integer,
    segment integer
        references segments
);

alter table user_segment
    owner to sivovgg;

