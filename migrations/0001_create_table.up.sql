-- FILE : 0001_create_table.up.sql

create table IF NOT EXISTS LOG_ENTRY(
    actor varchar(255) NOT NULL,
    action varchar(255) NOT NULL,
    occuredAt TIMESTAMPT NOT NULL
)