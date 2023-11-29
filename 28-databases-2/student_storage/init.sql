CREATE TABLE students (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS grades (
  id SERIAL PRIMARY KEY,
  student_id INT REFERENCES students(id),
  grade FLOAT
);
