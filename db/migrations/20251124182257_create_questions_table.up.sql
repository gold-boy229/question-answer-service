create table questions (
    id serial primary key,
    text text not null,
    created_at timestamptz default now() 
);
