create table answers (
    id serial primary key,
    question_id int not null,
    user_id text not null,
    text text not null,
    created_at timestamptz default now() 
);

alter table answers 
add constraint fk_answers_questions foreign key (question_id) references questions(id);
  