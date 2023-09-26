CREATE TABLE IF NOT EXISTS fio {
  id varchar not null unique primary key,
  name varchar not null, 
  surname varchar,
  age int,
  gender varchar,
  gender_probability float
}

CREATE TABLE IF NOT EXISTS fio {
  id varchar not null unique primary key,
  name varchar not null, 
  surname varchar,
  age int,
  gender varchar,
  gender_probability float
}